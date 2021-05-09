package main

import (
	"fmt"
)

func main() {
	for num1 := 1; num1 <= 10; num1++ {
		if num1 == 5 {
			fmt.Println("Matched, value =", num1)
			// After getting the desired value, the loop will exit. It won't continue.
			break
		} else {
			fmt.Println("Increased by 1 = ", num1)
		}
	}
}




/* 
__Output__:-
Increased by 1 =  1
Increased by 1 =  2
Increased by 1 =  3
Increased by 1 =  4
Matched, value = 5
*/
