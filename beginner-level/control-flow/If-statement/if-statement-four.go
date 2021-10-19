// https://play.golang.org/p/GXZK4LdoxK8

package main

import (
	"fmt"
)

var plot string
var pinCode int

func main() {
	// Identify the write plot and pinCode
	plot = "R-21"
	pinCode = 320012

	if plot == "R-21" {
		fmt.Println("PINCODE =", pinCode)
		fmt.Println("You're in correct plot.")
	} else {
		fmt.Println("Enter correct information.")
	}

}

/*
_Output_:-
PINCODE = 320012
You're in correct plot.

*/
