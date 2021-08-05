package main

import "fmt"

func main() {
	// main code block
	fmt.Println("Enter two values")
  var num1 int
	var num2 int
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	if num1 == num2 {
		fmt.Println("Both are same.")
	} else if num1 < 10 {
		fmt.Println("Number 1 is below 10.")
	} else if num2 < 10 {
		fmt.Println("Number 2 is below 10.")
	} else {
		fmt.Println("It's out of the number.")
	}
}
