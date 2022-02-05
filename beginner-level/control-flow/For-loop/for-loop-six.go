package main

import "fmt"

func main() {
	// Main code block
	// "for" loop with "continue"
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
		continue
	}
	fmt.Println("Task completed!!")

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
10
Task completed!!

*/
