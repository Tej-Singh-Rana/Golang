// capacity --> maximum size of the slice. || length --> size of the array.
// Never stored the data in the Slice. Slice data is always stored in an array.
// Slice is just a reference to the data. or the slice is just a reference to the array underneath.
// when we pass slices to functions, they automatically get past by reference.
// Slices work automatically like a pointers. slice > function > array > function > slice

package main

import "fmt"

func main(){
		// initialization 
		//make([]type, length, capacity)
		// at present it's value is 0. 
			count := make([]string, 5, 10)
			// fmt.Println(count)

			fmt.Printf("Length is: %d \nCapacity is: %d",len(count),cap(count))
		// ---------------------- or ---------------------------------------------
		// shorthand declaration of slice
			strike := []string{"slice1","slice2","slice3","slice4","slice5","slice6"}

			fmt.Printf("\nLength is: %d \nCapacity is: %d",len(strike),cap(strike))
		
		//----------------------index-value----------------
		// index value starts from zero.
		number :=[]int{150,120,450,200,670,2300,2500,1290,1550,9890}
		fmt.Println("\nindex number[3] value is:",number[3])
		
		//-----------------------manipulation-------------------
		number[3] = 1500
		fmt.Println("\nafter manipulate number[3] value:",number)	

		//---------------------slice-of-slice ------------------
		// slice[inclusive:exclusive], not include the exclusive value.
			sliceOfslice := number[2:8]
			fmt.Println("\nvalue of a new slice :",sliceOfslice[:5])

	}