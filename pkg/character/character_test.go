package character

import (
	"testing"
)

func TestIsLikeRule(t *testing.T) {
	result := IsLikeRule("a", "A")

	if !result {
		t.Errorf("Fail test like rule")
	}

	result = IsLikeRule("A", "a")

	if !result {
		t.Errorf("Fail test like rule")
	}

	result = IsLikeRule("ç", "Ç")

	if !result {
		t.Errorf("Fail test like rule")
	}
}

func BenchmarkIsLikeRule(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IsLikeRule("ç", "Ç")
	}
}

func TestNormalize(t *testing.T) {
	test := "á"

	want := "Á"

	result := Normalize(test, want)

	if result != want {
		t.Errorf("Fail normalize letter")
	}
}

func BenchmarkNormalize(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Normalize("ã", "Ã")
	}
}
