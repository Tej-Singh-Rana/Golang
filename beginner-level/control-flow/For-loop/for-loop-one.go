package main

import (
	"fmt"
)

var (
	count int
)

func main() {
	// Main func block
	count = 20
	for count := 20; count < 25; count++ {
		fmt.Println(count)
	}

}

/*
__Output__:-
20
21
22
23
24
*/


