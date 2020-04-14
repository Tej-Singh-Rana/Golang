// passing by value
// call by value
// In custom function, passing a copy of the value stored in some variable.
// Go passes by value rather than by reference.
package main

import "fmt"


func main(){
		name := "John"
		channel := "Netflix"

			fmt.Println(name,"watching movies on",channel + ".")

		// passing the variable channel value
	 // call the variable, passed it in our new variable
	 // copy of an argument and passes the copy, rather than original variable, 
	 // to the stack for the function being called. 
		changeChannel(channel)
		fmt.Println("\nWatching movies on",channel +".")
}

// func changeChannel(channel string)(choice string){

// 			channel = "Amazon Prime."
// 			choice = channel
// 			fmt.Println("\nTrying to watch the movies on",choice)

// 			return choice
// }
								// input parameter no need to be same.
								// (parameter func signature means -->what type we are passing in) return type --> what type will be returning.
func changeChannel(video string)string{
			// not declaring a new variable, so not used colon equal ":=" sign.
			// assigning a new value to an existing variable.
			video = "Amazon Prime."
			fmt.Println("\nTrying to watch the movies on",video)

			return video
}