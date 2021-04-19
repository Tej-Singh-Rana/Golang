package main // main package declaration

// Import package "fmt"
import (
	"fmt"
)

func main() {
	var s int
	for i := 1; i < 10; i++ {

		// variable "i" can be used within a brackets. Outside of the brackets it can't be used.
		s += s + i
	}

	// "fmt" package has Println() function.
	fmt.Println(s)

	// Trying to get a value of var "i"
	// ERROR - undefined i
	// fmt.Println(i)    // <-- uncomment that code, you will get an error as described above.
}
