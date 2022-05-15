// zero value comparison

package main

import "fmt"

func main() {

	// name1 := Name{First: "Rahul", Last: "Bourne"}   // 1st case
	name1 := Name{First: "", Last: ""} // 2nd case
	// it gives us a empty string.
	if name1 == (Name{}) {
		fmt.Println("Match")
	}

}

type Name struct {
	First string
	Last  string
}
