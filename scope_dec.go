package main

import "fmt"

// scope variables, define in outside function so this value can be taken by any function.
// { function blocks} inside this variables called local exists and outside curly brackets { } that's called global variables.
var comp string = "Global"

func main() {
		fmt.Println("Scope variables : ", comp)

}

func company() {
		fmt.Println("company is ", comp)

}
