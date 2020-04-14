// constants are immutable. once we assign them a value, we can't change it.


package main

import "fmt"

func main(){
	// we cannot assign in shorthand notation (":=").
	// assign by const keyword.
		const value = 1200
			fmt.Println("This is the constant value :",value)
		// value = 1250
		// fmt.Println("updated const value is ",value)  // got error.
}