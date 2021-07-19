package main

import (
	"fmt"
	c "runtime"
)

// var x int = 9
// var y int = 10

func main() {
	// code block
	// switch case

	// Declared and initialized variable x
	// After matching with the first case, it will check the if/else statement also.
	switch x := 9; x {
	// case - 1
	case 9:
		fmt.Println("System Info...")
		fmt.Println("OS type =", c.GOOS)
		if c.GOOS == "windows" {
			fmt.Println("Windows 10 Pro")
		} else if c.GOOS == "linux" {
			fmt.Println("Linux")
		}
	// case - 2
	case 10:
		fmt.Println("Linux")
	default:
		fmt.Println("Not Found")

	}
}

/*
Output :-
System Info...
OS type = linux
Linux
*/
