package iso7064

import (
	"testing"
)

func TestModulo97Radix10(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "32142829123456987654321611", want: "82"},
		{input: "794", want: "44"},
		{input: "1733", want: "40"},
		{input: "538182357", want: "82"},
		{input: "001937967935", want: "37"},
		{input: "352415823471", want: "14"},
		{input: "3214282912345698765432161182", want: "95"},
		{input: "100100100987654321131400", want: "85"},
		{input: "36155444216779151", want: "49"},
		{input: "77277287827223785", want: "90"},
		{input: "22181321400515123456741100", want: "07"},
		{input: "2218132140051512345674161100", want: "76"},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T) {
			if got := Modulo97Radix10(tt.input); got != tt.want {
				t.Errorf("Modulo97Radix10() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkModulo97Radix10(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		Modulo97Radix10("22181321400515123456741100")
	}
}
