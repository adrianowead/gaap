// Package sentence is specialist to tokenize given text on sentences
package sentence

import (
	"strings"
)

// SentenceTokenizer given string and split sentences
func SentenceTokenizer(text string) []string {
	var sentences = strings.Split(text, `\n`)

	return sentences
}
