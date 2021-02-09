package main

import (
	"fmt"
)

func main() {
	var i int // variable name type
	i = 42

	var j int = 23
	fmt.Printf("%v,%T", j, j)
	k := 55
	fmt.Println(i, j, k)
}
