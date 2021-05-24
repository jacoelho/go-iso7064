package iso7064

import (
	"testing"
)

func TestNormaliseNumeric(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "0123456789",
			want:  "0123456789",
		},
		{
			input: "012 345 678 9",
			want:  "0123456789",
		},
		{
			input: "A012-345-678-9C",
			want:  "0123456789",
		},
		{
			input: "ABCD",
			want:  "",
		},
		{
			input: "A012-345-678-9C 🍔",
			want:  "0123456789",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T) {
			if got := NormaliseNumeric(tt.input); got != tt.want {
				t.Errorf("NormaliseNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormaliseAlphabetic(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "0123456789",
			want:  "",
		},
		{
			input: "ABCD",
			want:  "ABCD",
		},
		{
			input: "abcd",
			want:  "ABCD",
		},
		{
			input: "a-b-c-d 🍔",
			want:  "ABCD",
		},
		{
			input: "abcd123",
			want:  "ABCD",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T) {
			if got := NormaliseAlphabetic(tt.input); got != tt.want {
				t.Errorf("NormaliseAlphabetic() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestNormaliseAlphaNumeric(t *testing.T) {
	tests := []struct {
		input string
		want  string
	}{
		{
			input: "0123456789",
			want:  "0123456789",
		},
		{
			input: "ABCD",
			want:  "ABCD",
		},
		{
			input: "abcd",
			want:  "ABCD",
		},
		{
			input: "a-b-c-d 🍔",
			want:  "ABCD",
		},
		{
			input: "abcd123",
			want:  "ABCD123",
		},
		{
			input: "aB1 c",
			want:  "AB1C",
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.input, func(t *testing.T) {
			if got := NormaliseAlphaNumeric(tt.input); got != tt.want {
				t.Errorf("NormaliseAlphaNumeric() = %v, want %v", got, tt.want)
			}
		})
	}
}
