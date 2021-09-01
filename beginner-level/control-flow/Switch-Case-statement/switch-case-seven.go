package main // declared main package

// imported fmt package
import (
	"fmt"
)

func main() {
	// main func body
	// switch case statement

	foo := '*'

	// Match the given value
	switch foo {
	case 0, 1, 2, 3, 4, 5:
		fmt.Println("Numeric case")
	case '*', ')', '(', '&', '+', '-':
		fmt.Println("Special case")
	default:
		fmt.Println("Not able to detect")

	}

}


/*
_Output_:-
Special case



*/
