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
	switch x := 9; x {
	case 9:
		fmt.Println("System Info...")
		fmt.Println("OS type =", c.GOOS)
		if c.GOOS == "windows" {
			fmt.Println("Windows 10 Pro")
		} else if c.GOOS == "linux" {
			fmt.Println("Linux")
		}
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
