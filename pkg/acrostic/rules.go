package acrostic

import (
	"regexp"
)

// Rules define the structure of current expression
type Rules struct {
	object string // {}
	count  int    // ||
	each   int    // []
	cmd    string // ^ $ #
	exp    string // ""
}

var basicStructure = "{}||[]"
var validCommands = []string{`^`, `$`, `#`}

// ParseRules process the received expression and returns Rules struct
func ParseRules(rules string) (bool, Rules) {
	out := Rules{}

	valid := HaveValidStructure(rules, &out)

	return valid, out
}

// HaveValidStructure validate if expression have minimum structure
func HaveValidStructure(rules string, r *Rules) bool {

	haveBasic := haveBasicStruct(rules, r)
	haveExp := haveBasicExp(rules, r)

	return haveBasic && haveExp
}

func extractStruct(rules string) string {
	re := regexp.MustCompile("[0-9a-zA-Z]")

	exp := re.ReplaceAllString(rules, "")

	return exp
}

func haveBasicStruct(rules string, r *Rules) bool {
	exp := extractStruct(rules)

	exp = exp[:len(basicStructure)]

	return exp == basicStructure
}

func haveBasicExp(rules string, r *Rules) bool {
	exp := extractStruct(rules)

	r.cmd = exp[len(basicStructure):][:1]

	valid := false

	for _, item := range validCommands {
		if item == r.cmd {
			valid = true
		}
	}

	return valid
}
