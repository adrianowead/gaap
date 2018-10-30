package acrostic

import (
	"fmt"
	"os"
)

// Error contains the position of error and strings needed
type Error struct {
	Received     string
	Needed       string
	TextPosition int
}

// IsValidRules returns true or false, if current rule is valid
// to acrostic method
func IsValidRules(rule string) (bool, Rules) {
	valid, r := ParseRules(rule)

	return valid, r
}

// ApplyRules receives te original text and rule to validate
// returs slices only with erros, and boolean to define if have errors or not
func ApplyRules(text string, rule string) (bool, []Error) {
	// validate rules
	valid, r := IsValidRules(rule)

	fmt.Println(valid)
	fmt.Println(r)
	os.Exit(1)

	return false, []Error{}
}
