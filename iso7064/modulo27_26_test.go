package iso7064

import (
	"testing"
)

func TestModulo27_26(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "JEJLMGJ", want: "S"},
		{input: "MUFEMSTCATLIT", want: "B"},
		{input: "VAQKBDHZQDYVZIATTNETJULCDAVRMQIEKIBD", want: "D"},
		{input: "OWNYDSZNWIBFVBRWRA", want: "U"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			if got := Modulo2726(tt.input); got != tt.want {
				t.Errorf("Modulo2726() = %v, want %v", got, tt.want)
			}
		})
	}
}
