package acrostic

import (
	"regexp"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// Status contains the position of error and strings needed
type Status struct {
	Received     string
	Needed       string
	TextPosition int
	isError      bool
}

// IsValidRules returns true or false, if current rule is valid
// to acrostic method
func IsValidRules(rule string) (bool, Rules) {
	valid, r := ParseRules(rule)

	return valid, r
}

// ApplyRules receives te original text and rule to validate
// returs slices only with erros, and boolean to define if have errors or not
func ApplyRules(text string, rule string) (bool, []Status) {
	// validate rules
	valid, r := IsValidRules(rule)

	isOk := false
	error := []Status{}

	if valid {
		// process sentence
		if r.Cmd == "@" {
			if r.Object == "P" {
				isOk, error = applySentenceParagraph(text, &r)
			}
		}
	}

	return isOk, error
}

func applySentenceParagraph(text string, r *Rules) (bool, []Status) {

	pagraph := strings.Split(text, "\n")

	for i := 0; i < len(pagraph); i++ {
		if len(strings.TrimSpace(string(pagraph[i]))) == 0 {
			// deleting
			pagraph = append(pagraph[:i], pagraph[i+1:]...)
		}
	}

	r.Exp = clearPunctuation(r.Exp)
	r.Exp = removeAccents(r.Exp)
	r.Exp = strings.TrimSpace(r.Exp)
	r.Exp = strings.Replace(r.Exp, " ", "", -1)

	validChars := regexp.MustCompile(`[^a-zA-Z0-9]`)

	posRule := 0
	out := []Status{}

	allOk := true
	beforeP := 0 - r.Each
	countOk := 0

	for k := 0; k < len(pagraph); k++ {

		currP := strings.TrimSpace(string(pagraph[k]))

		if processThisParagraph(beforeP, k, r, len(pagraph)) && len(currP) > 0 && len(r.Exp) > posRule {
			beforeP = k

			tmp := string([]rune(currP)[0])
			tmp = removeAccents(tmp)
			tmp = strings.ToLower(tmp)

			v := validChars.ReplaceAllString(tmp, "")
			curr := strings.ToLower(string(r.Exp[posRule]))

			o := Status{}
			o.Received = v
			o.Needed = curr
			o.TextPosition = k
			o.isError = true

			// only if is a valid char
			if len(v) > 0 && tmp == curr && curr != " " && curr != "" {
				o.isError = false
				countOk++
			}

			out = append(out, o)

			posRule++
		}
	}

	if countOk == len(r.Exp) {
		allOk = true
	}

	return allOk, out
}

func clearPunctuation(text string) string {
	re := regexp.MustCompile(`[^0-9a-zA-Z áàâãéèêíïóôõöúçñÁÀÂÃÉÈÍÏÓÔÕÖÚÇÑF]`)
	text = string(re.ReplaceAllString(text, " "))

	re = regexp.MustCompile(` +`) // remove double spaces
	text = string(re.ReplaceAllString(text, " "))

	return text
}

func removeAccents(text string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	s, _, _ := transform.String(t, text)

	return s
}

func processThisParagraph(before int, index int, r *Rules, countParagraph int) bool {
	valid := false
	newIdx := index + 1

	if len(r.EachOp) > 0 {
		if r.EachOp == "+" && newIdx == (before+r.Each) {
			valid = true
		} else if r.EachOp == "%" && index%r.Each == 0 {
			valid = true
		} else if r.EachOp == ":" && newIdx <= r.Each {
			valid = true
		}
	} else if len(r.EachOpAfter) > 0 {
		if r.EachOpAfter == ":" && countParagraph-r.Each <= newIdx {
			valid = true
		}
	}

	return valid
}
