// Package sentence is specialist to tokenize given text on sentences
package sentence

import (
	"strings"
)

// Tokenizer given string and split sentences
func Tokenizer(text string) []string {
	var sentences = strings.Split(text, "\n")
	var out []string

	for _, sentence := range sentences {
		var tmp = strings.Trim(sentence, " ")

		// discart empty lines
		if len(tmp) > 0 {
			out = append(out, tmp)
		}
	}

	return out
}
