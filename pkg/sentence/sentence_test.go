package sentence

import (
	"testing"
)

func TestTokenizer(t *testing.T) {
	test := "Mussum Ipsum, cacilds vidis litro abertis.\nMé faiz elementum girarzis, nisi eros vermeio.\nPraesent vel viverra nisi.\nMauris aliquet nunc non turpis scelerisque, eget.\nPaisis, filhis, espiritis santis. Interagi no mé, cursus quis, vehicula ac nisi."

	result := Tokenizer(test)

	// expected return
	want := []string{
		"Mussum Ipsum, cacilds vidis litro abertis.",
		"Mé faiz elementum girarzis, nisi eros vermeio.",
		"Praesent vel viverra nisi.",
		"Mauris aliquet nunc non turpis scelerisque, eget.",
		"Paisis, filhis, espiritis santis. Interagi no mé, cursus quis, vehicula ac nisi."}

	if len(result) != len(want) {
		t.Errorf("Length sentences results '%d', '%d' received", len(result), len(want))
	}

	for idx, text := range want {
		if text != result[idx] {
			t.Errorf("Position '%d' resturns ('%s'), expected ('%s')", idx, result[idx], text)
		}
	}
}

func BenchmarkTokenizer5sentences(b *testing.B) {
	test := "Mussum Ipsum, cacilds vidis litro abertis.\nMé faiz elementum girarzis, nisi eros vermeio.\nPraesent vel viverra nisi.\nMauris aliquet nunc non turpis scelerisque, eget.\nPaisis, filhis, espiritis santis. Interagi no mé, cursus quis, vehicula ac nisi."

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Tokenizer(test)
	}
}

func BenchmarkTokenizer2sentences(b *testing.B) {
	test := "Mussum Ipsum, cacilds vidis litro abertis.\nMé faiz elementum girarzis, nisi eros vermeio."

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		Tokenizer(test)
	}
}
