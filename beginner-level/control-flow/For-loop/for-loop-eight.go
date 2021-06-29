package main

import (
	"fmt"
)

const (
	tagOne  = "web"
	tag_Two = "db"
)

func main() {
	// Main code block
	print("Adding each tag to each group...\n")

	// var i value incremented by 1
	// var i value incremented by 2
	for i := 0; i < 3; i++ {
		fmt.Println(i) // 1st round --> 0, 2nd round --> 1, 3rd --> 2
		if i == 1 {
			if tagOne == "web" {
				// In 2nd round, i == 1 so it will print --> Web Server: web
				fmt.Println("Web Server:", tagOne)
			}
		}
		if i == 2 {
			if tag_Two == "db" {
				// In 3rd round, i == 2 so it will print --> Database Server: db
				fmt.Println("Database Server:", tag_Two)
			}
		}
		fmt.Println(i) // 1st round --> 0, 2nd round --> 1, 3rd round --> 2
	}
	// Printed the last value (2) and exited the loop.
}


/*
__Output__:-
Adding each tag to each group...
0
0
1
Web Server: web
1
2
Database Server: db
2

*/
