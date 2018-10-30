package acrostic

import (
	"regexp"
	"strconv"
	"strings"
)

// Rules define the structure of current expression
type Rules struct {
	Object      string // {}
	Count       int    // ||
	Each        int    // []
	EachOp      string // %(mod) : * +
	EachOpAfter string // :
	Cmd         string // ^ $ #
	Exp         string // ""
	Sub         *Rules
}

var basicStructure = "{}||[]"
var validCommands = []string{`^`, `$`, `#`}
var validObjects = []string{`P`, `W`}
var validEachOp = []string{`%`, `:`, `*`, `+`}
var validEachOpAfter = []string{`:`}

// ParseRules process the received expression and returns Rules struct
func ParseRules(rules string) (bool, Rules) {
	r := &Rules{}

	valid := parseCommand(rules, r)

	return valid, *r
}

func parseCommand(rules string, out *Rules) bool {
	valid := HaveValidStructure(rules, out)
	out.Exp = extractExpression(rules, out.Cmd)

	// if have subcommand
	if valid && out.Cmd == "#" && len(out.Exp) > 0 {
		out.Sub = &Rules{}

		out.Exp = out.Exp[1:]
		out.Exp = out.Exp[:len(out.Exp)-1]

		valid, out.Sub = parseSubCommand(out.Exp)
	}

	if valid {
		valid = extractObject(rules, out)
	}

	if valid {
		valid = extractCount(rules, out)
	}

	if valid {
		valid = extractEachCounter(rules, out)
	}

	return valid
}

// HaveValidStructure validate if expression have minimum structure
func HaveValidStructure(rules string, r *Rules) bool {

	haveBasic := haveBasicStruct(rules)
	haveExp := haveBasicExp(rules, r)

	return haveBasic && haveExp
}

func extractStruct(rules string) string {
	re := regexp.MustCompile(`[^[]{}|^#$]`)

	exp := re.ReplaceAllString(rules, "")

	re = regexp.MustCompile(`[0-9a-zA-Z:%]`)

	exp = re.ReplaceAllString(exp, "")

	return exp
}

func extractObject(rules string, r *Rules) bool {
	re := regexp.MustCompile(`\{(.*?)\}`)

	r.Object = re.FindString(rules)

	r.Object = r.Object[1:]
	r.Object = r.Object[:len(r.Object)-1]

	valid := false

	for _, item := range validObjects {
		if item == r.Object {
			valid = true
		}
	}

	return valid
}

func extractCount(rules string, r *Rules) bool {
	re := regexp.MustCompile(`\|(.*?)\|`)

	count := re.FindString(rules)

	re = regexp.MustCompile(`[0-9]`)
	count = re.FindString(count)

	i, err := strconv.Atoi(count)

	r.Count = i

	return err == nil
}

func extractEachCounter(rules string, r *Rules) bool {
	re := regexp.MustCompile(`\[(.*?)\]`)

	count := re.FindString(rules)

	count = count[1:]
	count = count[:len(count)-1]

	r.EachOp = string(count[0])

	valid := false

	for _, item := range validEachOp {
		if item == r.EachOp {
			valid = true
		}
	}

	if !valid {
		r.EachOp = ""
		r.EachOpAfter = string(count[1:])

		count = count[:1]
		count = count[len(count)-1:]

		for _, item := range validEachOpAfter {
			if item == r.EachOpAfter {
				valid = true
			}
		}
	}

	re = regexp.MustCompile(`[0-9]`)
	count = re.FindString(count)

	if len(r.EachOp) == 0 && len(r.EachOpAfter) == 0 && !valid {
		valid = true
	}

	i, err := strconv.Atoi(count)

	r.Each = i

	return err == nil && valid
}

func haveBasicStruct(rules string) bool {
	exp := extractStruct(rules)

	exp = exp[:len(basicStructure)]

	return exp == basicStructure
}

func haveBasicExp(rules string, r *Rules) bool {
	exp := extractStruct(rules)

	r.Cmd = exp[len(basicStructure):][:1]

	valid := false

	for _, item := range validCommands {
		if item == r.Cmd {
			valid = true
		}
	}

	return valid
}

func extractExpression(rules string, cmd string) string {
	out := ""

	i := strings.Index(rules, cmd)

	if i > -1 {
		out = rules[i+1:]
	}

	return out
}

func parseSubCommand(sub string) (bool, *Rules) {
	r := &Rules{}

	valid := parseCommand(sub, r)

	return valid, r
}
