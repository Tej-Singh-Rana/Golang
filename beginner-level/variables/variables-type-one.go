package main

import "fmt"

// Declared type only (int, string, bool etc..)

var (
	x int
	y int
	z string
)

func main() {
	x = 5
	y = 6
	z = "Declared only var type"
	fmt.Printf("Value of x and y: %d, %d\n", x, y)
	fmt.Printf("Comment: %s", z)
}