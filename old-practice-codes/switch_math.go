// In switch case statements, we can assign multiple comma separated values.
// 

package main

import (
		"fmt"
		"math/rand"   // importing module rand & time
		"time"
)

func main(){
			// declared random() func 
			switch tmpNum := random(); tmpNum {
		  // assign multiple values separated by comma
			case 0, 2, 4, 6, 8:
				fmt.Println("We got an even number :",tmpNum)
			case 1, 3, 5, 7, 9:
				fmt.Println("We got an odd number :",tmpNum)

			}
}			// random() func initialization 
		 //  func func-name(no parameter) type{}
			func random() int{
					rand.Seed(time.Now().Unix())   // 
					return rand.Intn(10)		// fixed the range of tmpNum
			}
