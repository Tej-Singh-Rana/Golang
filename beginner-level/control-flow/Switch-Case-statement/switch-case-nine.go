package main // declared main package

// imported fmt package
import (
	"fmt"
)

func Color(x string) string {
	return x

}

func main() {
	// Main func block

	// func Color() with switch case
	switch Color("red") {

	case "red":
		fmt.Println("Color is red.")
	case "black":
		fmt.Print("Color is black.\n")
	default:
		fmt.Println("Recheck...")
	}

}

/*
_Output_:-
Color is red.

*/
