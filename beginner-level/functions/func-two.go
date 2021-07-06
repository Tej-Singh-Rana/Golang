package main

import (
	"fmt"
)

// Multiple function declaration

func summerNation() string {

	// returning a string value and no argument passed

	return "Too hot"
}

func coldNation(wait string) string {

	return wait
}

func main() {

	// Initialize a value with short declaration

	vole := summerNation()

	fmt.Println("What?", vole)

	cole := coldNation("It's late")

	fmt.Println(cole)

}

/*
_Output_:-
What? Too hot
It's late

*/
