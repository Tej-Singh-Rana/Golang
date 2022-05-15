/* boolean true and false (or on and off)
boolean values &&(and), ||(or), !(not)
expression    	 value
true && true  	 true
true && false  	 false
false && true		 false
false && false   false
true || true     true
true || false    true
false || true    true
false || false   false
!true          	 false
!false           true
*/

package main

import ( 
			"fmt"
			"math"
)

func main()  {
			var x float64
			x = math.Pow(10,8)
			fmt.Println("number is : ", x-1)
			fmt.Println("Multiplication of two number : ", 32132 * 42452)
			var char = "Stay Safe and Be careful Go! Corona"
			fmt.Println("Length of string : ",len(char))	
			fmt.Println((true && false) || (false && true) || !(false && false))
}