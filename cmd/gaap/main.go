package main

import (
	"fmt"

	"../../pkg/acrostic"
	"../../pkg/tokenize"
)

func main() {

	//absPath, _ := filepath.Abs("testdata/lero-lero.txt")

	//text, err := ioutil.ReadFile(absPath)

	//check(err)

	//testSentence(string(text))
	testAcrosticRules(`{P}|1|[2]^"a"`)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func testSentence(text string) {
	var sentences = tokenize.TextToSentence(text)
	var words = tokenize.TextToWord(text)

	fmt.Printf("%d sentences, %d words, first word: '%s' %s", len(sentences), len(words), words[0], []byte{13, 10})
}

func testAcrosticRules(rules string) {
	fmt.Println(acrostic.IsValidRules(rules))
}
