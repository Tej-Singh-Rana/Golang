// Slices :-- A slice is formed by specifying two indices, a low and high bound, separated by a colon(:) format => a[low:high]
// low is the index of where to start the slice and high is the index where to end it (but not including the index itself).
// A slice is a segment of an array. Like arrays slices are indexable and have a length. 
// var x []float64 --> variable  variable-name [array-length]type.
// We can create a slice as well with the help of --> built-in function "make" function.
// format x := make([]float64,int,string...-type,length,capacity)
// copy and append function
package main

import "fmt"

func main(){
	// copy function used.	
	x := []int{ 1, 2, 3, 4, 5, 6}
	p := make([]int,1)
	copy(p,x)				// element x are copied into element p.
	fmt.Println(x,p)

}
