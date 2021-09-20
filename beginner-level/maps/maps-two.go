// https://play.golang.org/p/YBVSJtuEQKl

package main

import "fmt"

var m map[string]string

func main() {
	// Main code block
	// Map value is nil, as you can see in the output.
	fmt.Println(m)

	// Adding values
	// We can't assign key/value in the nil map
	m["test"] = "Fail"
	fmt.Println(m)

}

/*
_Output_:-
map[]
panic: assignment to entry in nil map

goroutine 1 [running]:
main.main()

*/
