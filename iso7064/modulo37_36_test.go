package iso7064

import (
	"testing"
)

func TestModulo37_36(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "A12425GABC1234002", want: "M"},
		{input: "RABD", want: "A"},
		{input: "X0H", want: "X"},
		{input: "B7Q3SFTUSH2QN7BIXBPMNZAM", want: "I"},
		{input: "UFSYYJO9FCID766EJYAEINTG43UZDD84MT6ZUDH08OM4N1K", want: "N"},
		{input: "ROYL38YZ9TDGPNB5MT40CEWGURAOKF07XEYTV3M", want: "6"},
		{input: "TBR", want: "1"},
		{input: "EFWW032G2TI", want: "U"},
		{input: "9", want: "J"},
		{input: "B3739U6CR", want: "K"},
		{input: "H0DJFUS8HHGZNEE9H6ZWW", want: "O"},
		{input: "C8AWF5G0CE8U9VTKSPPS2JAP09ZFEGFEAV", want: "L"},
		{input: "OW", want: "N"},
		{input: "TRO", want: "D"},
		{input: "1M", want: "Q"},
		{input: "FLFOWRIBFCNWNMNFVKAJVS7REUS2L", want: "2"},
		{input: "7LI6P5WTF2JHU", want: "U"},
		{input: "KUHNOF8OA1NXCW", want: "A"},
		{input: "KV0MFQXMAL4W5ICNH", want: "P"},
		{input: "MM61BF7H6C2O86NNMW8ZY8V", want: "8"},
		{input: "WP8Z0", want: "9"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			if got := Modulo3736(tt.input); got != tt.want {
				t.Errorf("Modulo3736() = %v, want %v", got, tt.want)
			}
		})
	}
}
