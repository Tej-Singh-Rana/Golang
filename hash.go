package main

import "fmt"

func main() {
	 const x string = "Hello World"
// x = "Global"     // cannot assign to x again because it's constant variable.
	fmt.Println(x)
	var (
		a = 10
		b = 15
		c = 20
	)
	fmt.Println("Value added ", a+b+c)
	const (
		car = "BMW"
		color = "Black"

	)
//	color := "Red"  // cannot assign to color again
	fmt.Println("car name : ", car)
	fmt.Println("Color is : ", color)
}
