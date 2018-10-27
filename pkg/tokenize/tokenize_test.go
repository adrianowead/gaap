package tokenize

import (
	"testing"
)

func TestTextToSentence(t *testing.T) {
	test := "Mussum Ipsum, cacilds vidis litro abertis.\nMé faiz elementum girarzis, nisi eros vermeio.\nPraesent vel viverra nisi.\nMauris aliquet nunc non turpis scelerisque, eget.\nPaisis, filhis, espiritis santis. Interagi no mé, cursus quis, vehicula ac nisi."

	result := TextToSentence(test)

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

func BenchmarkTextToSentence(b *testing.B) {
	test := "Mussum Ipsum, cacilds vidis litro abertis.\nMé faiz elementum girarzis, nisi eros vermeio.\nPraesent vel viverra nisi.\nMauris aliquet nunc non turpis scelerisque, eget.\nPaisis, filhis, espiritis santis. Interagi no mé, cursus quis, vehicula ac nisi."

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		TextToSentence(test)
	}
}

func TestStringToByte(t *testing.T) {
	test := "@3testando_ 😀 😁 😂 🤣 😃 😄 😅 😆 😉 😊 😋 😎 😍 😘, é nóis jão"

	want := []byte{
		64, 51, 116, 101, 115, 116, 97, 110, 100, 111, 95, 32, 240, 159,
		152, 128, 32, 240, 159, 152, 129, 32, 240, 159, 152, 130, 32,
		240, 159, 164, 163, 32, 240, 159, 152, 131, 32, 240, 159, 152,
		132, 32, 240, 159, 152, 133, 32, 240, 159, 152, 134, 32, 240,
		159, 152, 137, 32, 240, 159, 152, 138, 32, 240, 159, 152, 139,
		32, 240, 159, 152, 142, 32, 240, 159, 152, 141, 32, 240, 159,
		152, 152, 44, 32, 195, 169, 32, 110, 195, 179, 105, 115, 32,
		106, 195, 163, 111}

	result := StringToByte(test)

	if len(result) != len(want) {
		t.Errorf("Length bytes results '%d', '%d' received", len(result), len(want))
	}

	for idx, text := range want {
		if text != result[idx] {
			t.Errorf("Position '%d' resturns ('%d'), expected ('%d')", idx, result[idx], text)
		}
	}
}

func BenchmarkStringToByte(b *testing.B) {
	test := "@3testando_ 😀 😁 😂 🤣 😃 😄 😅 😆 😉 😊 😋 😎 😍 😘, é nóis jão"

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		StringToByte(test)
	}
}

func TestByteToString(t *testing.T) {
	test := []byte{
		64, 51, 116, 101, 115, 116, 97, 110, 100, 111, 95, 32, 240, 159,
		152, 128, 32, 240, 159, 152, 129, 32, 240, 159, 152, 130, 32,
		240, 159, 164, 163, 32, 240, 159, 152, 131, 32, 240, 159, 152,
		132, 32, 240, 159, 152, 133, 32, 240, 159, 152, 134, 32, 240,
		159, 152, 137, 32, 240, 159, 152, 138, 32, 240, 159, 152, 139,
		32, 240, 159, 152, 142, 32, 240, 159, 152, 141, 32, 240, 159,
		152, 152, 44, 32, 195, 169, 32, 110, 195, 179, 105, 115, 32,
		106, 195, 163, 111}

	want := "@3testando_ 😀 😁 😂 🤣 😃 😄 😅 😆 😉 😊 😋 😎 😍 😘, é nóis jão"

	result := ByteToString(test)

	if len(result) != len(want) {
		t.Errorf("Length bytes results '%d', '%d' received", len(result), len(want))
	}

	if result != want {
		t.Errorf("Received '%s', expected '%s'", result, want)
	}
}

func BenchmarkByteToString(b *testing.B) {
	test := []byte{
		64, 51, 116, 101, 115, 116, 97, 110, 100, 111, 95, 32, 240, 159,
		152, 128, 32, 240, 159, 152, 129, 32, 240, 159, 152, 130, 32,
		240, 159, 164, 163, 32, 240, 159, 152, 131, 32, 240, 159, 152,
		132, 32, 240, 159, 152, 133, 32, 240, 159, 152, 134, 32, 240,
		159, 152, 137, 32, 240, 159, 152, 138, 32, 240, 159, 152, 139,
		32, 240, 159, 152, 142, 32, 240, 159, 152, 141, 32, 240, 159,
		152, 152, 44, 32, 195, 169, 32, 110, 195, 179, 105, 115, 32,
		106, 195, 163, 111}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ByteToString(test)
	}
}

func TestNormalizeNewLine(t *testing.T) {
	test := []byte{13, 10}

	want := []byte{10}

	result := NormalizeNewLine(test)

	if string(result[:]) != string(want[:]) {
		t.Errorf("Falied to normalize new line")
	}
}

func BenchmarkNormalizeNewLine(b *testing.B) {
	test := []byte{13, 10}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		NormalizeNewLine(test)
	}
}
