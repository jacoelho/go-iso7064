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
		{input: "W1KDMNBFIZIVIDJQQ0F76S", want: "9P"},
		{input: "YJLDUW2XAT6JD346NRWT9Y", want: "KB"},
		{input: "R7TTSWFWIRP3PND1E42XUO", want: "VO"},
		{input: "87AVMBPFNQTY5RSQKSQ6JH", want: "92"},
		{input: "XVMZN7CD83796I1Q65VVZA", want: "0J"},
		{input: "0TXLE9L2FBO7ZZ1A6QWBBH", want: "6M"},
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
