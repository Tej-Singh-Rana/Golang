// comparing struct

package main

import "fmt"

func main() {

	name1 := Name{First: "Ram", Last: "Charan"}
	name2 := OtherName{First: "Ram", Last: "Charan"}

	if name1 == name2 {
		fmt.Println("Match")
	}

}

type Name struct {
	First string
	Last  string
}

// struct cannot match different-different types at same scenario even key:value is match.
// invalid operation: name1 != name2 (mismatched types Name and Other)
type OtherName struct {
	First string
	Last  string
}
