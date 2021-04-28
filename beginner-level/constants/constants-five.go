package main

import (
	"fmt"
)

// Package level declaration
// Short declaration(:=) only can use within that function block not an outside of that function. Ex:- foo := 20
// Can say that it's auto detect "type" according to an assignment.
const msg = "NoOne"

func main() {
	// cannot assign to msg (declared const)
	// 	msg = "NoBody"

	fmt.Println(msg)

	rowdy()
}

func rowdy() {
	// cannot assign to msg (declared const)
	// msg = "EveryBody"
	fmt.Println(msg)
}
