# Control Flow

#### Categories :-

 - For Loop
```go
package main

func main() {

	// Sample format of "For" loop
	// for <initialize>; <expression> {}

	for i := 0; i < 10; i++ {
		// code block
		print("Value of \"i\" = ", i, "\n")

	}
}
```

 - If/Else Statement
```go
package main

func main() {
	// Sample code "if/else" statement
	x := "sam"
	y := "john"

	if x != "dev" {
		print("No, it's not sam!!\n")
	} else if x == "sam" {
		print("Yes, it's sam!!\n")
	}

	if y != "john" {
		print("No, it's not john!!\n")
	} else if y == "john" {
		print("Yes, it's john!!\n")
	}
}
```

- Switch Statement
```go
package main

import "fmt"

func main() {
	// Sample code "switch" statement
	goo := 10
	switch goo {
	case 10:
		fmt.Println("Correct case.")
	case 20:
		fmt.Println("Incorrect case.")
	default:
		fmt.Println("None")
	}
}
```

