// Package tokenize given string on a sentence or word
package tokenize

import (
	"bytes"

	"../sentence"
)

// TextToSentence given a string and returns separeted senteces
func TextToSentence(text string) []string {
	b := StringToByte(text)
	b = NormalizeNewLine(b)

	text = ByteToString(b)

	var sentences = sentence.Tokenizer(text)

	return sentences
}

// StringToByte converts any string do byte slices
func StringToByte(text string) []byte {
	return []byte(text)
}

// ByteToString converts bytes to string
func ByteToString(b []byte) string {
	return string(b[:])
}

// NormalizeNewLine convert any endlines to linux endline
func NormalizeNewLine(d []byte) []byte {
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)

	// replace CF \r (mac) with LF \n (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)

	return d
}
