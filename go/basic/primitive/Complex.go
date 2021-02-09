package main

import (
	"fmt"
)

func main() {
	var a complex128 = 1 - 3i
	b := 3 + 30i
	fmt.Printf("%v,%T", a, a)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
}
