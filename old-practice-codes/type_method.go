// In Go, initialize variable value is zero
// package level variables are global
// Shorthand declare and initialize ":=" only works inside of functions.
// Variables declared at the function level must be used.

package main           // entry point

import (
			 "fmt"
			 "reflect"
)

var (
			name, module string = "John", "reflect"
			number       int = 1450
)
	// var  name string = "john"
	// var module string
	// var number int

func main(){

			fmt.Println("Name  is : ",name," and is type ",reflect.TypeOf(name))
			fmt.Println("Module  is : ",module," and is type : ",reflect.TypeOf(module))
			fmt.Println("Number is : ",number," and is type : ",reflect.TypeOf(number))
}

