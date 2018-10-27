// Package tokenize given string on a sentence or word
package tokenize

import (
	"bytes"

	"../sentence"
)

// TextToSentence given a string and returns separeted senteces
func TextToSentence(text string) []string {
	b := StringToByte(text)
	b = normalizeNewLine(b)

	text = byteToString(b)

	var sentences = sentence.Tokenizer(text)

	return sentences
}

// StringToByte converts any string do byte slices
func StringToByte(text string) []byte {
	return []byte(text)
}

func byteToString(b []byte) string {
	return string(b[:])
}

func normalizeNewLine(d []byte) []byte {
	// replace CR LF \r\n (windows) with LF \n (unix)
	d = bytes.Replace(d, []byte{13, 10}, []byte{10}, -1)

	// replace CF \r (mac) with LF \n (unix)
	d = bytes.Replace(d, []byte{13}, []byte{10}, -1)

	return d
}
