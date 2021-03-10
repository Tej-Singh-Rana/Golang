package main // package declaration 

var (
	x int = 15 // global variable, defined outside of a function. It can be use anywhere in the program.
	y int = 20
)

// func body
func cup(x int) int {
	// It will give an error.
	// print(tab)
	return x
}

func main() {
	tab := 12 // local variable, declared inside a function. It can be use by within a function or block where its declared.
	print("Total of x+y = ", x+y, "\n")
	z := cup(25)
	print("Value of x, z and tab ", x, ", ", z, ", ", tab, "!!")

}

