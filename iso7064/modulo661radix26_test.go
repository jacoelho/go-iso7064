package iso7064

import (
	"testing"
)

func TestModulo661Radix26(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "XKFSHTWWCOMYYASPSYTHJW", want: "CJ"},
		{input: "BAISDLAFK", want: "BM"},
		{input: "SOMETHING", want: "IF"},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T) {
			if got := Modulo661Radix26(tt.input); got != tt.want {
				t.Errorf("Modulo661Radix26() = %v, want %v", got, tt.want)
			}
		})
	}
}
