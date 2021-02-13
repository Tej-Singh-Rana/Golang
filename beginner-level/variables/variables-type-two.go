package main

import "fmt"

// If we have one or more variable types that are the same then we can define in a single line. 
var (
	road, trade string
	distance    int
	count       bool
)

func main() {
	road = "Lions"
	trade = "coin trade"
	distance = 1200
	count = true
	fmt.Printf("%s is a %s and %d is %v value.", road, trade, distance, count)

}