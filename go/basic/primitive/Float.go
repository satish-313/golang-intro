package main

import (
	"fmt"
)

func main() {
	n := 3.14
	fmt.Printf("%v,%T\n", n, n)
	n = 13.7e45
	fmt.Printf("%v,%T\n", n, n)
	n = 2.1e14
	fmt.Printf("%v,%T\n", n, n)
}
