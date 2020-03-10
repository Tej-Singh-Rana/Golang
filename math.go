package main

import (
		"fmt"
		"math"
)
func main()  {
	var h1 float64 = 132.89
	smile := '😀'
	value := math.Sqrt(h1)
	fmt.Printf("Square root of %v is %.2f!!",h1,value)
	fmt.Printf("\nthis is the complete square root value of %v is %g!!",h1,value)
	fmt.Printf("\n%d %f %t %c %q %v %#U", smile,smile,smile,smile,smile,smile,smile)
	
	per := math.Round(value)
	set := math.Ceil(value)
	fo := math.Floor(value)
	fmt.Printf("\nRound figure value of %v is %.2f ",value,per)
	fmt.Printf("\nCeil value is %.2f",set)     //It's goes up to round figure  
	fmt.Printf("\nFloor value is %.2f",fo)    //It's goes down to round figure
}