package main

import (
	"fmt"
)

// on package level we can't declare colon syntax we have to use full declare
var (
	i         int    = 50
	actorName string = "salman shit"
)

func main() {
	fmt.Println(i)
	// fmt.Printf("%v,%T", i, i) // format string it sucks
	var i int = 22 // in our case inner var take presedent is called shodowing
	// i := 33.14
	k := "satish"
	fmt.Println(i, actorName, "king:", k)
}
