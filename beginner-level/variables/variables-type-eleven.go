package main // package declaration

import "fmt" // import 'fmt' package

var (
	x = 12
)

// func body
func main() {

	// multiple short-hand var declaration
	// scope of this var, only in this curly brackets. We cant use it's value globally (in other functions).
	x, y, z := 4500, "plus", 45.56

	fmt.Println(x, y, z)

	Demo()
}

func Demo() int {
	p, _ := fmt.Println("Value of global x variable =", x)

	return p

}

/*
_Output_:-
4500 plus 45.56
Value of global x variable = 12

*/
