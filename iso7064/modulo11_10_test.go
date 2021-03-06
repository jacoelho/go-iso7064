package iso7064

import (
	"testing"
)

func TestModulo11_10(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "079", want: "2"},
		{input: "0794", want: "5"},
		{input: "6511224300430482", want: "7"},
		{input: "53266878780357001264776215687313785128868", want: "4"},
		{input: "003", want: "2"},
		{input: "7846306445607555258161616166", want: "2"},
		{input: "861543053353545631188671371628487317133258177413", want: "0"},
		{input: "4555773880", want: "8"},
		{input: "05320162033", want: "8"},
		{input: "71065204127183146616272676714658212652753", want: "4"},
		{input: "4716", want: "6"},
		{input: "68058451423857734108017363773364282850032", want: "1"},
		{input: "4674678664653", want: "1"},
		{input: "42328026", want: "0"},
		{input: "500300224343584662684600506004225567445126724117", want: "5"},
		{input: "65428072048156805364127530485864402163255", want: "9"},
		{input: "308655184386", want: "3"},
		{input: "4317754156343048540758528678637685134", want: "4"},
		{input: "4600034236774", want: "1"},
		{input: "1385812104511745558672550524138724810", want: "9"},
		{input: "1563104756864506624875436082452546036618466424", want: "6"},
		{input: "71660618558507745431171460", want: "5"},
		{input: "3676012560274058125725682603573122447174631135", want: "4"},
		{input: "853", want: "7"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			if got := Modulo1110(tt.input); got != tt.want {
				t.Errorf("Modulo1110() = %v, want %v", got, tt.want)
			}
		})
	}
}
