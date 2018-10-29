package character

import (
	"strings"
	"unicode"
)

// IsLikeRule returns boolean to test if current text character is equal current character rule
func IsLikeRule(characterText string, characterRule string) bool {
	equal := false

	characterText = Normalize(characterText, characterRule)

	// convert characterText to rule
	if characterText == characterRule {
		equal = true
	}

	return equal
}

// Normalize convert first character to equal second (upper to lower, or lower to upper)
func Normalize(a string, b string) string {
	tmpA := []rune(a)[0]
	tmpB := []rune(b)[0]

	out := a

	if unicode.IsUpper(tmpA) && unicode.IsLower(tmpB) {
		out = strings.ToLower(a)
	}

	if unicode.IsLower(tmpA) && unicode.IsUpper(tmpB) {
		out = strings.ToUpper(a)
	}

	return out
}
