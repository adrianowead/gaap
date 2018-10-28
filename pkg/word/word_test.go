package word

import (
	"testing"
)

func TestTextToWord(t *testing.T) {
	test := "gaap é o melhor    software em    GoLang!"

	want := []string{"gaap", "é", "o", "melhor", "software", "em", "GoLang!"}

	result := TextToWord(test)

	if len(result) != len(want) {
		t.Errorf("Length bytes results '%d', '%d' received", len(result), len(want))
	}

	for idx, text := range want {
		if text != result[idx] {
			t.Errorf("Position '%d' resturns ('%s'), expected ('%s')", idx, result[idx], text)
		}
	}
}

func BenchmarkTextToWord(b *testing.B) {
	test := "gaap é    o melhor software  em GoLang!"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TextToWord(test)
	}
}
