package main

import "fmt"

func main()  {

	sum := 1

	for sum < 500 {
		sum +=sum
		fmt.Println(sum)
	}
    fmt.Println(sum)	
}