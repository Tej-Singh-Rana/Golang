// adding key-value pair in map function
// also updating old few keys value

package main

import "fmt"

func	main(){

	x := map[uint]string{
		1: "int",
		2: "str",
		3: "float",
		4: "list",
		5: "dict",
	}
	fmt.Println("Old list :-- ",x)
	// adding new key-value
	x[0] = "variable"
	x[6] = "tuple"
	x[7] = "function"
	x[8] = "int/str"
	fmt.Println("New added value updated :-- ",x)
	// updating old keys values
	
	 x[0] = "var"
	 x[1] =  "data-type"
	 x[2] = "method"
	 fmt.Println("Some old keys value changes :-- ",x)
}