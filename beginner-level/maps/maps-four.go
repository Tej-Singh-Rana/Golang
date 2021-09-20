// https://play.golang.org/p/vNz3VZWXFyd

package main

import "fmt"

var key map[string]string

func main() {
	// Main code block
	// Declare map "key" value
	fmt.Println(key)
	key = map[string]string{"marvel": "avengers"}
	fmt.Println(key)

	// Initializing more key/value
	key["warner"] = "bros"
	fmt.Println(key)

	// Declare and initialize new map
	// with make function
	diffTask := make(map[int]string)
	fmt.Println("-----------------------------------------------")
	fmt.Println("with make function")
	fmt.Printf("Type of new map = %T, Value = %v\n", diffTask, diffTask)

	diffTask[1] = "Good"
	diffTask[2] = "Bad"
	fmt.Println("Reviews", diffTask)
}

/*
_Output_:-
map[]
map[marvel:avengers]
map[marvel:avengers warner:bros]
-----------------------------------------------
with make function
Type of new map = map[int]string, Value = map[]
Reviews map[1:Good 2:Bad]

*/
