package iso7064_test

import (
	"fmt"

	"github.com/jacoelho/go-iso7064/iso7064"
)

func ExampleNormaliseNumeric() {
	fmt.Println(iso7064.NormaliseNumeric("1A3 4"))
	// Output: 134
}

func ExampleNormaliseAlphabetic() {
	fmt.Println(iso7064.NormaliseAlphabetic("aB1c"))
	// Output: ABC
}

func ExampleNormaliseAlphaNumeric() {
	fmt.Println(iso7064.NormaliseAlphaNumeric("aB1 c"))
	// Output: AB1C
}
