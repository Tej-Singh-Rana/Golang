// Arguments evaluated at time defer is executed, not at time of called function execution.

package main

import "fmt"

func main(){
		a := "defer"
		fmt.Println("What will be print?")
		defer fmt.Println(a)

		a = "function"
		// defer fmt.Println(a)
}