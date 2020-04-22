//
package main

import (
	"errors"
	"fmt"
)

func main() {

	port := 5000
	_, err := service(port)
	fmt.Println("Server is started on port ", port)
	fmt.Println("error : ", err)
}

func service(port int) (int, error) {

	fmt.Println("Starting server...")
	fmt.Println("Server is going to process...")
	x := errors.New("Something went wrong!!")
	return port, x
}
