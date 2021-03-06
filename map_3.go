// Retrieving values
// Modifiying data
// Using delete function

package main

import "fmt"

func main(){
// adding some key value pairs
var user = map[string]string{
"User1":	"Red color",
"User2":	"Brownish color",
"User3":	"White wine",
"User4":	"Dry martini with a twist",
"User5":	"Top sales performer",
}
fmt.Print("Data :-- ",user)
fmt.Print("\n")
// Retrieving values

val1 := user["User3"]
val2 := user["User5"]

fmt.Println("User3 data is : ",val1)
fmt.Println("User5 data is : ",val2)

// Modifiying data

newUser := user

newUser["User6"] = "Investors to the world"
newUser["User7"] = "Marketing meetings schedule"

fmt.Println("User's and their data :-- ", newUser)

// Deleting 
// delete function format :- delete(map-name,key)
delete(newUser,"User3")
delete(newUser,"User6")

fmt.Println("After deleted, updated list :-- ", newUser)
// checking user existing or not
// boolean :- true || false
t1, ok := newUser["User1"]
if ok == true{
	fmt.Println("User is present.")
	fmt.Println("User data is : ",t1)
}else {
	fmt.Println("User doesn't exist.")
}

}