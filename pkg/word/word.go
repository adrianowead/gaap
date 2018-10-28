package word

import (
	"strings"
)

// TextToWord receives a string and return words
func TextToWord(text string) []string {
	res := strings.Split(text, " ")

	var out []string

	for _, el := range res {
		tmp := strings.TrimSpace(el)

		// discart empty
		if len(tmp) > 0 {
			out = append(out, tmp)
		}
	}

	return out
}
