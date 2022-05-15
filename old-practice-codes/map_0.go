// map function

package main

import "fmt"

func main(){

	// var x map[int]string
	// x[1] = "value"

	// fmt.Println(x)
// we will get an error because map value is nil so we cannot add any key-value pair.
// so build a nil map through make function

var x = make(map[int]string)
x[1] = "value"

fmt.Println(x)
fmt.Println(x[2])   // map returns the zero value which element doesn't exist, because of empty string data.
// method value, boolean condition
// Accessing an element of a map can return two values instead of just one. 
// The first value is the result of the lookup, the second tells us whether or not the lookup was successful in true||false.

name, ok := x[1]
fmt.Println("1. ",name,"2. ",ok)
}