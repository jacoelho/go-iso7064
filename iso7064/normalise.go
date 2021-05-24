package iso7064

import (
	"strings"
)

// copied from utf8 package
// characters below runeSelf are represented as themselves in a single byte.
const runeSelf = 0x80

func isNumeric(b uint8) bool {
	return '0' <= b && b <= '9'
}

func isAlphabeticUppercase(b uint8) bool {
	return 'A' <= b && b <= 'Z'
}

func isAlphabeticLowercase(b uint8) bool {
	return 'a' <= b && b <= 'z'
}

func isAlphabetic(b uint8) bool {
	return isAlphabeticUppercase(b) || isAlphabeticLowercase(b)
}
func isAlphaNumeric(b uint8) bool {
	return isNumeric(b) || isAlphabeticUppercase(b) || isAlphabeticLowercase(b)
}

func alphabeticLowerToUpper(b uint8) uint8 {
	return b - 'a' + 'A'
}

func identity(c uint8) uint8 { return c }

func convertToUpper(c uint8) uint8 {
	if !isAlphabeticLowercase(c) {
		return c
	}

	return alphabeticLowerToUpper(c)
}

func normalise(predicate func(uint8) bool, convert func(uint8) uint8, input string) string {
	sb := new(strings.Builder)

	for i := 0; i < len(input); i++ {
		c := input[i]

		if c >= runeSelf {
			continue
		}

		if predicate(c) {
			sb.WriteByte(convert(c))
		}
	}

	return sb.String()
}

// NormaliseNumeric filters a string ignoring non digit characters
// "1A3 4" is converted to "134"
func NormaliseNumeric(input string) string {
	return normalise(isNumeric, identity, input)
}

// NormaliseAlphabetic filters a string ignoring non ascii letters characters
// lower case ascii letters are converted to uppercase
// "aB1c" is converted to "ABC"
func NormaliseAlphabetic(input string) string {
	return normalise(isAlphabetic, convertToUpper, input)
}

// NormaliseAlphaNumeric filters a string ignoring non ascii letters or digit characters
// lower case ascii letters are converted to uppercase
// "aB1 c" is converted to "AB1C"
func NormaliseAlphaNumeric(input string) string {
	return normalise(isAlphaNumeric, convertToUpper, input)
}
