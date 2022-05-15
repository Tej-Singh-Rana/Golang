//A switch statement is a shorter way to write a sequence of if - else statements. 
//It runs the first case whose value is equal to the condition expression.

package main

import (
	"fmt"
	"time"
)
func main()  {
	const week = "Tuesday"
	switch week {
	case "Monday":								//case condition like "if"
		fmt.Println("Today is Monday.")
	case "Tuesday":
		fmt.Println("Today is Tuesday.")
	case "Wednesday":
		fmt.Println("Today is Wednesday.")
	case "Thursday":
		fmt.Println("Today is Thursday.")
	default:								//It's like "else" 
		fmt.Println("This is rest of weeks.")
	}

	fmt.Println("When's Wednesday?")
	today := time.Now().Weekday()
	switch time.Sunday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
		fmt.Println(today + 2)    //It will count next 2 days.
	}
	fmt.Println(time.Now())
}