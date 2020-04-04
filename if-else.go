package main

import "fmt"

func main()  {
	code := 2 
	// an "if" statement, it has a condition followed by a block. if the condition evaluates to true then it run block conditions.
	// an "if" statement is false then it will move to next "else if" or "else" condition if not present then it exit from block.
	// condition checked followed by top down.
	// the first one condition is true then it will not execute other blocks, even if their conditions also pass/true.
	if  code <= 5 {
			fmt.Println("This is less then or equal to 5")
	} else if code == 10 {
	 		 fmt.Println("This is more then value of code")
	} else if code > 5 {
			fmt.Println("This is greather then 5")
	} else {												//We have to write else statement after ending of }.
		 	fmt.Println("Try again for success")
}	
		fmt.Println("Value of code will print here ",code)
	}