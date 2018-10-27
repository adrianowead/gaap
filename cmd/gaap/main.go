package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"../../pkg/tokenize"
)

func main() {

	absPath, _ := filepath.Abs("testdata/lero-lero.txt")

	text, err := ioutil.ReadFile(absPath)

	check(err)

	testSentence(string(text))
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func testSentence(text string) {
	var tokens = tokenize.TextToSentence(text)

	// fmt.Println(tokens[0])
	fmt.Println(len(tokens))
}
