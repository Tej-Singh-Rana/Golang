package main

import (
	"fmt"
)

func main() {
	// Main code block
	// Declaring a pointer for variable "p" and type "string".
	var p *string

	// Variable is not holding any value.
	fmt.Println("Value not assigned for pointer 'p': ", *p)
}

/*
_Output_:-
panic: runtime error: invalid memory address or nil pointer dereference
[signal 0xc0000005 code=0x0 addr=0x0 pc=0xd9514a]

goroutine 1 [running]:
main.main()
        C:/Users/User/go/src/github.com/tej-singh-rana/Golang/beginner-level/pointers/pointers-one.go:13 +0x2a
exit status 2

*/
