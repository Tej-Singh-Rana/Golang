package main

import "fmt"

func main() {
		fmt.Print("Enter the value of Fahrenheit : ")
		var x float64 
		fmt.Scanf("%f", &x)
  // converting fahrenheit to celsius
		Cel := (x - 32) * 5/9
		fmt.Println("Fahrenheit to Celsius", Cel)
}

