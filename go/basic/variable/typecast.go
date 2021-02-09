package main

import (
	"fmt"
	"strconv"
)

func main() {
	var i int = 33
	fmt.Printf("%v,%T\n", i, i)
	// var j float32
	// j = float32(i)
	var j string
	j = strconv.Itoa(i) // in case of string it give ascii or unicode value instead of typecast
	fmt.Printf("%v,%T\n", j, j)
}
