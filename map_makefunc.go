// Map function with make function
// format of make function make([key-type]value-type,length,capacity)
// while using make function in map then make(map[key-type]value-type,capacity)
package main

import "fmt"

func main()  {
		// initializing 
		var x = make(map[int]string)
		// var x  map[int]string
		// if x == nil{
		// 	fmt.Println("True")
		// }else {
		// 	fmt.Println("False")
		// }
		fmt.Println(x)
		fmt.Println(len(x))

		// adding two values
		x[1] = "Base"
		x[2] = "Two"
		fmt.Println("Element one is : ", x[1])
		fmt.Println("Element two is : ", x[2])
}