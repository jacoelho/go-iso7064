package iso7064

// hybrid system recursive method
func hybrid(modulo int, from fromAlphabet, to toAlphabet, input string) string {
	sum := modulo
	for _, r := range input {
		intermediate := (sum + from(r)) % modulo

		if intermediate == 0 {
			intermediate = modulo
		}

		sum = (intermediate * 2) % (modulo + 1)
	}

	checksum := (1 + modulo - sum) % modulo

	return to(checksum)
}

// Modulo1110 generates check character in accordance with ISO/IEC 7064, MOD 11,10
// Designation 6
func Modulo1110(input string) string {
	return hybrid(10, numericToValue, numericToString, input)
}

// Modulo2726 generates check character in accordance with ISO/IEC 7064, MOD 27,26
// Designation 7
func Modulo2726(input string) string {
	return hybrid(26, alphaToValue, alphaToString, input)
}

// Modulo3736 generates check character in accordance with ISO/IEC 7064, MOD 37,36
// Designation 8
func Modulo3736(input string) string {
	return hybrid(36, alphanumericToValue, alphanumericToString, input)
}
