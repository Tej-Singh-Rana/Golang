// returning multiple values from a function using a return statement.
// we can define return values name as well but by default Golang automatically detect it.
package main

import "fmt"

// assigining multiple values
								   // this (int, int, int) part shows that function will return 3 int values
func f(a,b,c int) (int, int, int) {
	return a * b * c, a + b + c, a - b - c
}
// defines return value by names
func r(l1,l2 int) (val1, val2 int) {   //(val1 int, val2 int) 
	 val1  = l1*2
	 val2  = l2*5
	return	// val1, val2
}
func main(){
		x, y, z := f(25,10,30)

		fmt.Println("Value of multiply ",x)
		fmt.Println("Value of addition ",y)	
		fmt.Println("Value of subtraction ",z)
		
		t1, t2 := r(30,25)

		fmt.Println("Value of t1(val1) ", t1)
		fmt.Println("Value of t2(val2) ", t2)

		// blank identifier (_) underscore
		_, ok := r(t1,t2)
		fmt.Println(ok)
	}