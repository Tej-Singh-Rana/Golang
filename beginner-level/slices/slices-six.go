// https://play.golang.org/p/Yvkmew5XNhV

package main

import (
	"fmt"
)

// Declare variable newTag type []string
var newTag []string

func main() {
	// Main code block

	// Declare newTag type slice
	newTag = []string{"Shang-Chi", "Venom", "Spider-Man"}
	fmt.Println(newTag)

	// Append function
	newTag = append(newTag, "Avengers")
	fmt.Print("----------------------------------\n")
	fmt.Println("After appending new element into the slice: ", newTag)

	// We can add more elements in append function
	newTag = append(newTag, "Nobody", "Hulk", "Candyman")
	fmt.Print("----------------------------------\n")
	fmt.Println("After appending multiple elements into the slice: ", newTag)
}

/*
_Output_:-
[Shang-Chi Venom Spider-Man]
----------------------------------
After appending new element into the slice:  [Shang-Chi Venom Spider-Man Avengers]
----------------------------------
After appending multiple elements into the slice:  [Shang-Chi Venom Spider-Man Avengers Nobody Hulk Candyman]


*/
