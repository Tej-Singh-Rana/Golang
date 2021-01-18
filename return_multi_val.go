// multiple return values in go

package main

import "fmt"

func Name() (string, string) {
	return "First", "Last"
}

func call() (first string, last string) {
	first = "Tej Singh"
	last = "Rana"
	return
}

func main() {
	n1, n2 := Name()
	fmt.Println(n1, n2)

	m1, m2 := call()
	fmt.Println(m1, m2)
}
