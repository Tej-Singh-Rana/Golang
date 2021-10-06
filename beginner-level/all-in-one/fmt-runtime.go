package main

import (
	"fmt"

	// short alias
	os "runtime"
)

func main() {
	// Main code block

	fmt.Printf("GO OS : %v\n", os.GOOS)
	fmt.Printf("GO ARCH: %v\n", os.GOARCH)
	fmt.Printf("GO VERSION: %v\n", os.Version())
}

/*
_Output_:-
GO OS : linux
GO ARCH: amd64
GO VERSION: go1.17.1

*/
