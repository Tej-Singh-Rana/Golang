package main

import "fmt"

func main()  {
// variable assign  variable name [size of array]type
	var x [7]float64
	x[0] = 98
  x[1] = 93
  x[2] = 77
  x[3] = 82
  x[4] = 83
	x[5] = 100
	x[6] = 120
	var total float64 = 0
	// len(x) is an int, so we have to convert into float.
	// invalid operation: total / len(x) (mismatched types float64 and int)
	for i := 0; i < len(x); i++ {
  		total += x[i]
}
		fmt.Println(total / float64(len(x)))
	
}