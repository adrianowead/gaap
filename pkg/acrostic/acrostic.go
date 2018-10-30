package acrostic

// IsValidRules returns true or false, if current rule is valid
// to acrostic method
func IsValidRules(rule string) (bool, Rules) {
	valid, r := ParseRules(rule)

	return valid, r
}
