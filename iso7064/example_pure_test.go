package iso7064_test

import (
	"fmt"

	"github.com/jacoelho/go-iso7064/iso7064"
)

func ExampleModulo11Radix2() {
	// Modulo11Radix2
	fmt.Println(iso7064.Modulo11Radix2("82703512512057"))
	// Output: X
}

func ExampleModulo37Radix2() {
	// Modulo37Radix2
	fmt.Println(iso7064.Modulo37Radix2("A999915000001"))
	// Output: M
}

func ExampleModulo97Radix10() {
	// Modulo97Radix10
	fmt.Println(iso7064.Modulo97Radix10("36155444216779151"))
	// Output: 49
}

func ExampleModulo661Radix26() {
	// Modulo661Radix26
	fmt.Println(iso7064.Modulo661Radix26("BAISDLAFK"))
	// Output: BM
}

func ExampleModulo1271Radix36() {
	// Modulo1271Radix36
	fmt.Println(iso7064.Modulo1271Radix36("ISO79"))
	// Output: 3W
}
