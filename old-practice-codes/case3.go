package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	file, err := ioutil.ReadFile("F:\\python_online.txt")
	if err != nil {
		return
	}
	fmt.Println(string(file))
}
