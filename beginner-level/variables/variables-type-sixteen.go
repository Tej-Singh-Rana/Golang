package main // main package

// imported "fmt" package
import (
	"fmt"
)

func main() {

	{ // start block

		// Variable scope in inside of the block
		var train string = "self"
		var days int = 30
		fmt.Printf("Training method = %s and total number of days = %d\n", train, days)

	} // end block

	// Variable scope outside of the block, result is different compare to above block of code.
	var train string = "coach"

	fmt.Println("Printed \"train\" value (Training method) from outside of the block =", train)
}

