// Built-in types: arrays, slices and maps.
// Arrays :- An array is a numbered sequence of elements of a single type with a fixed length. x[5], x[10]
// arrays are indexed starting from 0. 

package main

import "fmt"

// func main(){
// 			var x[10] int
// 			x[4] = 100
// 			fmt.Println(x)

// }

func main(){
			var x[5] float64
			x[0] = 98
			x[1] = 93
			x[2] = 77
			x[3] = 82
			x[4] = 83

			var total float64 = 0
			for i := 0; i < 5; i++ {
				total += x[i]
				// fmt.Println(total)    print the total value by sequence
				// i = i + 1   increment by 1 of variable "i", till 4 and at 5, it exit.
			}
			fmt.Println(total / 5)

}