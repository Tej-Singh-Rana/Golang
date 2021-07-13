package main

import (
	"fmt"
)

func main() {
	// Main Code block
	// var name [length]type

	var weeks [7]string
	weeks[0] = "Sunday"
	weeks[1] = "Monday"
	weeks[2] = "Tuesday"
	weeks[3] = "Wednesday"
	weeks[4] = "Thursday"
	weeks[5] = "Friday"
	weeks[6] = "Saturday"

	fmt.Println(weeks)
	// Would like to print one by one.
	// Used underscore(_), bcuz we don't need that value at the moment.

	for _, week := range weeks {
		fmt.Println(week)
	}
}

/*
_Output_:-
[Sunday Monday Tuesday Wednesday Thursday Friday Saturday]
Sunday
Monday
Tuesday
Wednesday
Thursday
Friday
Saturday

*/

