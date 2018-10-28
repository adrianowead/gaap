// Package tokenize given string on a sentence or word
package tokenize

import (
	"bytes"

	"../sentence"
	"../word"
)

// TextToSentence given a string and returns separeted senteces
func TextToSentence(text string) []string {
	text = normalizeString(text)

	var sentences = sentence.Tokenizer(text)

	return sentences
}

func normalizeString(text string) string {
	b := StringToByte(text)
	b = NormalizeNewLine(b)

	return ByteToString(b)
}

// TextToWord receives a string and return words
func TextToWord(text string) []string {
	text = normalizeString(text)

	return word.TextToWord(text)
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
