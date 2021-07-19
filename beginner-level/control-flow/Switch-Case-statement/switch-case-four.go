package main // declared main package

// imported fmt package
import (
	"fmt"
)

func main() {
	// main func body
	// switch case statement
	var (
		mustang string = "red"
	)
	// Declared and initialized earlier. Now, we will match the value directly to the case.
	switch {
	case mustang == "black":
		fmt.Println("Color not matched!!")

	case mustang == "red":
		fmt.Println("Color matched and that's", mustang, ".")

	default:
		fmt.Println("Not Found!!")
	}

}

/*
_Output_:-
Color matched and that's red .

*/
