// variadic function
// when you don't know about parameters quantity 
// when you want to pass a slice in a function
// By using ... before the type name of the last parameter 
// , you can indicate that it takes zero or more of those parameters.
// increase the readability of program
package main

import "fmt"
// variadic function ...
// func func-name(parameter ...type)type
func add(args ...int)int{
		total := 0
		for _, v := range args{
			total += v
		}
		return total
}
func main(){

	fmt.Println(add(1,2,3,4,5,6)) // variadic function
	// initializing in slice function
	// [key-type]value-type{value1, value2...} 
	value := []int{200,500,1500,2000}
	fmt.Println(add(value...))
	// fmt.Println(value[0:3])
}