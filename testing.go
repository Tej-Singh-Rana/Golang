package main

import "fmt"

func main()  {

	// var name string
	// fmt.Println("Enter your name : ")
	// fmt.Scanf("%s",&name)
	// if name == ""{
	// 	fmt.Println("Name is not given.")
	// } else if name == "code"{
	// 		fmt.Println("You are verified.")
	// }else {
	// 		fmt.Println("Enter a valid name.")
	// }

	x := []float64{ 50, 56, 45, 78, 100, 12, 13, 111, 1222, 1232, 123231, 1234324, 12325435, 67567, 657, 5677}
	x[5] = 500
	fmt.Println(x[1:5])
}