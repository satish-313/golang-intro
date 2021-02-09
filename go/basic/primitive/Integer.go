package main

import (
	"fmt"
)

func main() {
	// n := 42
	var n uint16 = 43 // unsign interger
	fmt.Printf("%v,%T\n", n, n)

	a := 10             // 1010
	b := 3              // 0011
	fmt.Println(a & b)  // 0010
	fmt.Println(a | b)  // 1011
	fmt.Println(a ^ b)  // 1001
	fmt.Println(a &^ b) // 0100

	fmt.Println("math")
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
}
