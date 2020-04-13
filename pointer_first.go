// Initializing pointer

package main

import "fmt"

func main(){
			// normal declaration
			var a int
			a = 1500
			//declaration of pointer
			var p *int

			// initialization of pointer
			p = &a

			fmt.Println("Value of \"a\" ",a)
			fmt.Println("memory address of \"a\" : ", &a)
			fmt.Println("value stored in variable p : ",p)
		// value updated
			a = 2000
		 // initialization of pointer
			p = &a
			// assigning a new value to pointer
			*p = 4500

			fmt.Println("updated value of \"a\" : ",a)
			fmt.Println("updated value of \"a\" : ",*p)

}