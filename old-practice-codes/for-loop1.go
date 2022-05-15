package main

import "fmt"

func main()  {
	sum := 1
//Iteration
	for ; sum < 100 ; {

		sum += sum
		fmt.Println(sum)
	}
	fmt.Println(sum)
}