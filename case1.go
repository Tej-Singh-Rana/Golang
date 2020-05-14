package main

import (
	"fmt"
	"os"
)

func main() {

	file, err := os.Open("F:\\python_online.txt")
	if err != nil {
		return
	}
	defer file.Close()

	size, err := file.Stat()
	if err != nil {
		return
	}
	fmt.Printf("The file is %d bytes.", size.Size())
}
