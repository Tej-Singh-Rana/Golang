//
//

package main

import "fmt"

func main(){

		mySlice := make([]int, 1, 4)
		fmt.Printf("Length is: %d, Capacity is: %d",
							len(mySlice),cap(mySlice))
		// increase the size of capacity
		/*	capacity of the slice is equal to the length of its underlying array.
		 		it'll add values without increasing the size of the underlying array.
		 		once we hit the size of that underlying array or slice, 
				 add a fifth value, append is then going to double the size of the underlying array.
				 create a new array with twice the length.
				 array doubled from 4 to 8, 8 to 16 and then from 16 to 32.
		 */
		for i := 1; i < 17 ; i++{
			mySlice = append(mySlice, i)
			fmt.Printf("\nCapacity is: %d",
								cap(mySlice))
		}
		
}