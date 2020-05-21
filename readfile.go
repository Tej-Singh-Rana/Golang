package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func rat() {
	tube, err := ioutil.ReadFile("C:\\goplaybookcli\\TestPage.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File content is %s", tube[:])
}
func main() {
	rat()
}
