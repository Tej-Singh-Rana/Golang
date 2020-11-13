package main

import "fmt"

func main() {

	var cal int = 0  // variable defined

	for count := 0; count < 5; count++ {       // count++
			 cal += count         
			fmt.Println(cal) 
	}
	fmt.Println(cal)   // last value out
}

