package iso7064

import (
	"testing"
)

func TestModulo37Radix2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "G123489654321", want: "Y"},
		{input: "A999915000001", want: "M"},
		{input: "LMOSBYFCTRSSJCB3X6U9DUBLKNIH03FO6X963TGXYAHG", want: "*"},
		{input: "G7N286AL39FFQNU6BH8SMJB5VGMBCB6WPGMWJK800VU7K", want: "P"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			if got := Modulo37Radix2(tt.input); got != tt.want {
				t.Errorf("Modulo37Radix2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkModulo37Radix2(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		Modulo37Radix2("G123489654321")
	}
}
