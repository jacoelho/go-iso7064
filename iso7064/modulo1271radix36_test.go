package iso7064

import (
	"testing"
)

func TestModulo1271Radix36(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "SOMETHING", want: "3X"},
		{input: "ISO79", want: "3W"},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T) {
			if got := Modulo1271Radix36(tt.input); got != tt.want {
				t.Errorf("Modulo1271Radix36() = %v, want %v", got, tt.want)
			}
		})
	}
}
