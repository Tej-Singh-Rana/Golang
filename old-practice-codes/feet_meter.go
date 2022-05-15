package main

import "fmt"

func main(){
		fmt.Print("Enter value to convert feet into meters : ")
		var feet float64			// assign variable
		fmt.Scanf("%f", &feet)			// taking input 
// converting value 
		meter := feet * 0.3048			//  1 ft = 0.3048 m
		fmt.Println("Converted value in meter : ", meter)

}