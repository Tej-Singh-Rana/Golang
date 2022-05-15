// Functions are also known as method, sub-routine, or procedure.
// Collectively the parameters and the return type are known as the function's signature.
// stack, variadic, closure, recursion, defer, panic ,recover essential points.
package main

import "fmt"

func main(){
	xs := []float64{98,37,23,45,56}

	total := 0.0
	// for loop
	for i:=0; i < len(xs);i++ {
		total += xs[i]
		fmt.Println(xs[i])
	}
	fmt.Println(total / float64(len(xs)))
	
}
func average(xs []float64) float64{
		// fmt.Println(xs(5,6)))
		// built-in function panic to handle run-time error
		panic("Not Implemented")
}