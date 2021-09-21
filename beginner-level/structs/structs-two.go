// https://play.golang.org/p/rSGxjKu4Xre

package main

import (
	"fmt"
)

type Class struct {
	name  string
	email string
}

func main() {

	user := Class{"Rock", "rock@sample.com"}

	// %+v gives more details
	fmt.Printf("Lists of user: %+v\n", user)

	// Fetching individual values
	fmt.Println("User name: ", user.name)
	fmt.Println("Email address: ", user.email)
}

/*
_Output_:-
Lists of user: {name:Rock email:rock@sample.com}
User name:  Rock
Email address:  rock@sample.com


*/
