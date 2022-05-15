// A pointer holds the memory address of a value.
// The asterisk(*) operator denotes the pointer's underlying value.
// The memory address is always found in hexadecimal format(0x...) like 0xc0034fd.
// The default value or zero-value of a pointer is always nil.
// Uninitialized pointer will always have a nil value.
// "new function" takes a type as an argument, allocates enough memory to fit a value of that type and returns a pointer to it.
// format :-->  var pointer_name *data_type

package main

import (

		"fmt"
)

func main(){
				// using "new" func with pointer 
				// *string --> A pointer of type string which can store only the memory addresses of string variables.
				var firstName *string = new(string)
			// dereference operator as *(asterisk) used
			// point to firstName
				*firstName = "Fire stripe"
				// dereference operator
				fmt.Println(*firstName)

				var lastName = "Rapid fire"
				fmt.Println(lastName)
			// memory address 
				ptr := &lastName
				fmt.Println("value stored in pointer variable ptr : ",ptr)
				fmt.Println("Value of lastName :",*ptr)
			// memory address both have same memory address.
				lastName = "Rock and role"
				fmt.Println("value stored in pointer variable ptr : ",ptr)
				fmt.Println("Value of lastName :",*ptr)
}