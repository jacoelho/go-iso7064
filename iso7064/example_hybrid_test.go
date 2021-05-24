package iso7064_test

import (
	"fmt"

	"github.com/jacoelho/go-iso7064/iso7064"
)

func ExampleModulo1110() {
	// Modulo1110
	fmt.Println(iso7064.Modulo1110("4600034236774"))
	// Output: 1
}

func ExampleModulo2726() {
	// Modulo2726
	fmt.Println(iso7064.Modulo2726("MUFEMSTCATLIT"))
	// Output: B
}

func ExampleModulo3736() {
	// Modulo3736
	fmt.Println(iso7064.Modulo3736("B3739U6CR"))
	// Output: K
}
