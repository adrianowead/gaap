package acrostic

import (
	"regexp"
	"strings"
)

// Rules define the structure of current expression
type Rules struct {
	Object string // {}
	Count  int    // ||
	Each   int    // []
	EachOp string // %(mod) : * +
	Cmd    string // ^ $ #
	Exp    string // ""
	Sub    *Rules
}

var basicStructure = "{}||[]"
var validCommands = []string{`^`, `$`, `#`}

// ParseRules process the received expression and returns Rules struct
func ParseRules(rules string) (bool, Rules) {
	r := &Rules{}

	valid := parseCommand(rules, r)

	return valid, *r
}

func parseCommand(rules string, out *Rules) bool {
	valid := HaveValidStructure(rules, out)
	out.Exp = extractEspression(rules, out.Cmd)

	// if have subcommand
	if valid && out.Cmd == "#" && len(out.Exp) > 0 {
		out.Sub = &Rules{}

		out.Exp = out.Exp[1:]
		out.Exp = out.Exp[:len(out.Exp)-1]

		_, out.Sub = parseSubCommand(out.Exp)
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

func extractEspression(rules string, cmd string) string {
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
