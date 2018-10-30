package acrostic

// IsValidRules returns true or false, if current rule is valid
// to acrostic method
func IsValidRules(rule string) (bool, Rules) {
	valid, basicStruct := ParseRules(rule)

	return valid, basicStruct
}
