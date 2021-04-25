package main

import "fmt"

/* We use "const" keyword for constants. Ex:- const Key = "900"
Constants cannot be declared using the ":=" syntax. Like const Jago := "Alert"
We can use boolean, numeric, string and character in place of values. (keyword <variable-name> type = value)
Ex:-
const val bool = true
const strong string = "Power"
const speed int = 250
*/

const key int = 900

func main() {
	// In case, if you'll assign a new value again 1000 to "key".
	//** cannot assign to key (declared const) **
	// key = 1000
	fmt.Println(key)
}
