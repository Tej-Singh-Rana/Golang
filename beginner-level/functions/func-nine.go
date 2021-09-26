// https://play.golang.org/p/dvU8LdQmGcg

package main

import (
	"fmt"
)

func main() {
	// Main code block

	passMulti := arbitFunc(1, 2, 3, 4, 5)
	fmt.Println("Total number : ", passMulti)

}

// We can pass multiple inputs
// three dot (...)

func arbitFunc(numbers ...int) int {
	number := 0

	for _, num := range numbers {
		number += num

		fmt.Println(number)

	}

	return number

}

/*
_Output_:-
1
3
6
10
15
Total number :  15

*/
