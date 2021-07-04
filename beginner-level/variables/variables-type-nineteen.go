package main

import (
	"fmt"
)

// We can use backtick(`) instead of single or double quotes. Ex: foo = `most interesting variable in the programming world.`, boo = `Good`
// Same method applicable on "const".

var team = `Frozen Battle`

const year = `2021`

func main() {
	fmt.Println("Name of team:", team)
	fmt.Printf("Year: %v, %T", year, year)
}

/*
_Output_:-
Name of team: Frozen Battle
Year: 2021, string

*/
