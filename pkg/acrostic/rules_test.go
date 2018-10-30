package acrostic

import "testing"

func TestParseRules(t *testing.T) {
	test := []string{`{P}|1|[2]^"a"`,
		`{P}|1|[:2]^"a"`,
		`{P}|0|[%2]#({W}|0|[:2]^"br")`,
		`{P}|0|[2]#({W}|0|[2:]$"brazil")`,
		`{B}|1|[@2]^"a"`,
		`{P}|1|[:2]9"a"`}

	want := []bool{true, true, true, true, false, false}

	for idx, item := range test {
		result, _ := ParseRules(item)

		if result != want[idx] {
			t.Errorf("Fail to valid rule (%s), expected (%t), received (%t)", item, want[idx], result)
		}
	}
}

func BenchmarkParseRules(b *testing.B) {
	test := `{P}|0|[2]#({W}|0|[2:]$"brazil")`

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		ParseRules(test)
	}
}

func TestHaveValidStructure(t *testing.T) {
	test := []string{`{P}|1|[2]^"a"`,
		`{P}|1|[:2]^"a"`,
		`{P}|0|[%2]#({W}|0|[:2]^"br")`,
		`{P}|0|[2]#({W}|0|[2:]$"brazil")`,
		`{B}|1|[@2]^"a"`,
		`{P}|1|[:2]9"a"`}

	want := []bool{true, true, true, true, false, false}

	for idx, item := range test {
		r := &Rules{}

		result := HaveValidStructure(item, r)

		if result != want[idx] {
			t.Errorf("Fail to valid rule (%s), expected (%t), received (%t)", item, want[idx], result)
		}
	}
}

func BenchmarkHaveValidStructure(b *testing.B) {
	test := `{P}|0|[2]#({W}|0|[2:]$"brazil")`

	r := &Rules{}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		HaveValidStructure(test, r)
	}
}
