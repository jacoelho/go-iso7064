package iso7064

const (
	digits               = "0123456789"
	digitsPlusX          = digits + "X"
	alpha                = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphanumeric         = digits + alpha
	alphanumericPlusStar = digits + alpha + "*"
)

func numericToValue(r rune) int {
	return int(r - '0')
}

func numericToString(v int) string {
	return digits[v : v+1]
}

func numericPlusXToValue(r rune) int {
	if r == 'X' {
		return 10
	}
	return int(r - '0')
}

func numericPlusXToString(v int) string {
	return digitsPlusX[v : v+1]
}

func alphaToValue(r rune) int {
	return int(r - 'A')
}

func alphaToString(v int) string {
	return alpha[v : v+1]
}

func alphanumericToValue(r rune) int {
	switch {
	case r == '*':
		return 36

	case r >= '0' && r <= '9':
		return int(r - '0')

	default:
		return int(10 + r - 'A')
	}
}

func alphanumericToString(v int) string {
	return alphanumericPlusStar[v : v+1]
}
