// Package tokenize given string on a sentence or word
package tokenize

import (
	"../sentence"
)

// TextToSentence given a string and returns separeted senteces
func TextToSentence(text string) []string {
	b := StringToByte(text)

	return sentences
}

// StringToByte converts any string do byte slices
func StringToByte(text string) []byte {
	return []byte(text)
}
