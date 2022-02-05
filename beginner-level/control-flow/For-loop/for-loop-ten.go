package main

import "fmt"

func main() {
	// Main code block
	x := 0
	value := 10
	for {
		if x == value {
			break
		}
		fmt.Println(x)
		x++

	}
}

/*
_Output_:-
0
1
2
3
4
5
6
7
8
9

*/
