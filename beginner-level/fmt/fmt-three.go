package main

// Provide a shorter alternative name to a library's package name

import (
	shortName "fmt"
)

func main() {
	shortName.Println("It's a short alias name. Easy and effective.\n" +
   								"This line imports the fmt package but assigns it an alias shortName.\n" +
									"You can create an alias for an imported package to simplify its usage, especially if the original package name is long or if you want to avoid naming conflicts.\n" +
	                "shortName now serves as a shorthand for referring to the fmt package.")
}

/*
_Output_:-

It's a short alias name. Easy and effective.
This line imports the fmt package but assigns it an alias shortName.
You can create an alias for an imported package to simplify its usage, especially if the original package name is long or if you want to avoid naming conflicts.
shortName now serves as a shorthand for referring to the fmt package.

*/
