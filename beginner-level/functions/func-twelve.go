// https://go.dev/play/p/aE8C6WPda9a
package main

// Importing other packages to use their functions
import (
	"fmt"
	"strings"
)

func main() {
	// Main code block
	batchName := greet("coDex")
	fmt.Println(batchName)
}

func greet(batch string) string {
	// Simple function called greet function with keyword "func" and returing a value type string
	return strings.ToUpper(batch)
}

/*
_Output_:-
CODEX


*/
