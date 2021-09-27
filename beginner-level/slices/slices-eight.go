package main

import (
	"fmt"
)

func main() {
	// Main code block

	list := []string{"Rahul", "Rakesh", "Tej", "Garima", "Yukti", "Ram", "Himank", "Yash", "Lokesh", "Faruq"}

	// 0. Lenght
	fmt.Println(len(list))
	// 1. Print the last element
	fmt.Println(list[len(list)-1:])
	// 2. Print the element from 1 to 9
	fmt.Println(list[1:])
	// 3. Print the element from 3 to 5, not including 6th element
	fmt.Println(list[3:6])
	// 4. Print the element from 1 to 9
	fmt.Println(list[1 : len(list)-1])

}

/*
_Output_:-
10
[Faruq]
[Rakesh Tej Garima Yukti Ram Himank Yash Lokesh Faruq]
[Garima Yukti Ram]
[Rakesh Tej Garima Yukti Ram Himank Yash Lokesh]

*/
