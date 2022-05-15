package main

import "fmt"
func math(x int, y int)(int, int, int){
	var v1 = x + y
	var v2 = x - y
	var v3 = x * y
	return v1, v2, v3
}
func main(){
		c1 := 10
		c2 := 20
		result1,result2,result3 := math(c1,c2)
		fmt.Println("Result1 value is ",result1,"\nResult2 value is ",result2,"\nResult3 value is ",result3)
}
