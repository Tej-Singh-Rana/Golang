package main

import "fmt"

func Loop(x int,y int) (int, int) {
	var value1 = x + y
	var value2 = x - y
	return	value1, value2
}

func main(){
	num1 := 20
	num2 := 10
	var result1,result2 = Loop(num1,num2)
	fmt.Println(result1,result2)
	if result1 == 30 {
		fmt.Println("Value of num1 : ",result1)
	} else if result2 == 10{
		fmt.Println("Value of num2 : ", result2)
	} else {
			fmt.Println("Try it more")
	}
}