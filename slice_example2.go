// add slices as elements within a slice. 
//

package main

import "fmt"

func main(){

	mySlice := []int{1, 2, 3, 4, 5}
	fmt.Println(mySlice)

	for _, i := range mySlice{
			fmt.Println("for range loop :",i)

	}
	newSlice := []int{15,20,25,30,35,40}
	// slice to slice with the ellipses
	// instead of appending the slice, appends the element of the slice
	// appending a collection of integers to slice 
	mySlice = append(mySlice, newSlice...)
	fmt.Println(mySlice)

}