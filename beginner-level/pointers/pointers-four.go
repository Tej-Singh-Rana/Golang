// https://play.golang.org/p/lyEohUKW4pK

package main

import (
	"fmt"
)

func main() {
	// Main code block

	// integer pointer
	var cash *int

	widCash := 5000
	cash = &widCash

	fmt.Println(widCash, cash, *cash)

	// It's type
	fmt.Printf("cash type is: %T, widCash type is: %T", cash, widCash)
}

/*
_Output_:-
5000 0xc000112000 5000
cash type is: *int, widCash type is: int

*/
