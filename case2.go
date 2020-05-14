package main

import (
	"os"
)

func main() {

	file, err := os.Create("F:\\test.txt")
	if err != nil {
		return
	}
	defer file.Close()

