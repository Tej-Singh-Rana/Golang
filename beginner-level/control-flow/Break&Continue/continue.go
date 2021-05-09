package main

import (
	"fmt"
)

func main() {
	for num1 := 5; num1 <= 10; num1++ {
		if num1 == 6 {
			fmt.Println("Matched, value =", num1)
			// After getting the desired value, the loop won't exit. It'll continue till value 10.
			continue
		} else {
			fmt.Println("Increased by 1 = ", num1)
		}
	}
}



/* 
__Output__:-
Increased by 1 =  5
Matched, value = 6
Increased by 1 =  7
Increased by 1 =  8
Increased by 1 =  9
Increased by 1 =  10
*/
