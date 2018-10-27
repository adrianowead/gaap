// Package tokenize given string on a sentence or word
package tokenize

import (
	"../sentence"
)

// TextToSentence given a string and returns separeted senteces
func TextToSentence(text string) []string {
	var sentences = sentence.SentenceTokenizer(text)

	return sentences
}
