# https://play.golang.org/p/sBs6bxii4h2
package main

import "fmt"

// school type struct has name and class fields.
type school struct {
	name  string
	class int
}

func main() {
	fmt.Println(school{"Tej", 12})

	// Accessing struct fields with dot
	x := school{"John", 10}
	
	fmt.Println(x.name)
	fmt.Println(x.class)

}

/*
_Output_:-
{Tej 12}
John
10

*/
