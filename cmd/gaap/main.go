package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"../../pkg/acrostic"
	"../../pkg/tokenize"
)

func main() {

	// absPath, _ := filepath.Abs("testdata/lero-lero.txt")

	//text, err := ioutil.ReadFile(absPath)

	//check(err)

	//testSentence(string(text))
	// testAcrosticRules(`{P}|1|[2]^"a"`)
	// testAcrosticRules(`{P}|0|[%2]#({W}|0|[:2]^"br")`)
	// testAcrosticRules(`{P}|1|[:2]^"a"`)
	testAcrosticRules(`{P}|1|[:2]@"veras que um filho teu nao foge a luta"`)

	absPath, _ := filepath.Abs("testdata/sentence-acrostic.txt")
	text, _ := ioutil.ReadFile(absPath)

	testAcrosticSentence(string(text), `{P}|1|[%2]@"verás que um filho teu não foge a luta"`)
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
	// valid, out := acrostic.IsValidRules(rules)

	// fmt.Println(out, valid)
	// fmt.Println("")
	// fmt.Println(out.Sub)
}

func testAcrosticSentence(text string, rule string) {
	valid, proc, objectType := acrostic.ApplyRules(text, rule)

	fmt.Println(proc)
	fmt.Println(objectType)
	fmt.Println(valid)
}
