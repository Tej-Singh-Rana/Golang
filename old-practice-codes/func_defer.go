// It's schedules a function call to be run after the function completes.
// Deferred functions are run even if a run-time panic occurs.
// Used to delay execution of a statement until function exits.
// Useful to group "open" and "close" function together.
// Run in LIFO(Last in, First out) order
// Arguments evaluated at time defer is executed, not at time of called function execution.

package main

import ( 
			"fmt"
)

func main(){
			fmt.Println("Start")
			fmt.Println("Process")
			fmt.Println("End")
		fmt.Println("_____________________________________________")
		fmt.Println("Difference b/w normal func & defer func ")
			fmt.Println("First")
		defer fmt.Println("Second")   // defer function used, it will be execute after func exits.
		fmt.Println("Third")
		fmt.Println("_____________________________________________")
		fmt.Println("Difference b/w normal func & defer func ")
	defer fmt.Println("First Entry")
	defer fmt.Println("Second Entry")
	defer fmt.Println("Third Entry")
	fmt.Println("_____________________________________________")

}