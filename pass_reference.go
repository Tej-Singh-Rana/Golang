// & --> References a pointer
// * --> De-references a pointer
// Pointer && passing by reference

package main

import "fmt"


func main(){
		name := "John"
		channel := "Netflix"

			fmt.Println(name,"watching movies on",channel + ".")

		// passing the variable channel value
	 // call the variable, passed it in our new variable
	 // ampersand & --> sends to a variable to the pointer in memory location.
		changeChannel(&channel)
		fmt.Println("\nWatching movies on",channel +".")
		fmt.Println("\nLatest movies collection is on the way in",channel +".")
}
								// input parameter no need to be same.
								// *string --> video is pointer to string.
func changeChannel(video *string)string{
			// not declaring a new variable, so not used colon equal ":=" sign.
			// assigning a new value to an existing variable.
			// this value assigning into the memory that channel pointer is referencing (&channel).
			*video = "Amazon Prime"
			fmt.Println("\nTrying to watch the movies on",*video +".")

			return *video
}