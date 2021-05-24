package iso7064

import (
	"testing"
)

func TestModulo11Radix2(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{input: "000000021825009", want: "7"},
		{input: "000000029079593", want: "X"},
		{input: "079", want: "X"},
		{input: "0794", want: "0"},
		{input: "001175717748247", want: "6"},
		{input: "747633", want: "6"},
		{input: "734404529805608", want: "0"},
		{input: "41812925", want: "9"},
		{input: "986596515101003", want: "X"},
		{input: "177804390", want: "4"},
		{input: "899952", want: "6"},
		{input: "8200107028943", want: "6"},
		{input: "6583491086272", want: "6"},
		{input: "118344415502", want: "1"},
		{input: "982546", want: "4"},
		{input: "000194585", want: "9"},
		{input: "82703512512057", want: "X"},
		{input: "4338", want: "6"},
		{input: "007763161148", want: "8"},
		{input: "790830359455", want: "0"},
		{input: "4731722270", want: "0"},
		{input: "000088486619034", want: "0"},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.input, func(t *testing.T) {
			if got := Modulo11Radix2(tt.input); got != tt.want {
				t.Errorf("Modulo11Radix2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkModulo11Radix2(b *testing.B) {
	b.ReportAllocs()

	for n := 0; n < b.N; n++ {
		Modulo11Radix2("000000029079593")
	}
}
