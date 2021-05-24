package iso7064

type (
	// convert an alphabet character to a value
	fromAlphabet func(r rune) int

	// convert a digit to the alphabet character
	toAlphabet func(int) string
)

// pure system recursive method
// multiples of the modulo are discarded
func pure(modulus, radix int, doubleChecksum bool, from fromAlphabet, to toAlphabet, input string) string {
	var sum int
	for _, r := range input {
		v := from(r)
		intermediate := (sum + v) * radix
		sum = intermediate % modulus // discard modulus multiples
	}

	if doubleChecksum {
		sum = (sum * radix) % modulus
	}

	checksum := (1 + modulus - sum) % modulus

	if !doubleChecksum {
		return to(checksum)
	}

	return to(checksum/radix) + to(checksum%radix)
}

// Modulo11Radix2 generates check character in accordance with ISO/IEC 7064, MOD 11–2
// Designation 1
func Modulo11Radix2(input string) string {
	return pure(11, 2, false, numericPlusXToValue, numericPlusXToString, input)
}

// Modulo37Radix2 generates check character in accordance with ISO/IEC 7064, MOD 37–2
// Designation 2
func Modulo37Radix2(input string) string {
	return pure(37, 2, false, alphanumericToValue, alphanumericToString, input)
}

// Modulo97Radix10 generates check characters in accordance with ISO/IEC 7064, MOD 97–10
// Designation 3
func Modulo97Radix10(input string) string {
	return pure(97, 10, true, numericToValue, numericToString, input)
}

// Modulo661Radix26 generates check characters in accordance with ISO/IEC 7064, MOD 661–26
// Designation 4
func Modulo661Radix26(input string) string {
	return pure(661, 26, true, alphaToValue, alphaToString, input)
}

// Modulo1271Radix36 generates check characters in accordance with ISO/IEC 7064, MOD 1271–36
// Designation 5
func Modulo1271Radix36(input string) string {
	return pure(1271, 36, true, alphanumericToValue, alphanumericToString, input)
}
