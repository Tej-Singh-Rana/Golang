package main

import "fmt"

func role(name string) {
	// Assign role func, name -> string
	fmt.Printf("Admin Name: %s\n", name)
}

func main() {
	// calling role func in main body
	role("Tej Singh Rana")

	var Name string
  // highlight message
	fmt.Println("name of the user : ")
  // for user's input
	fmt.Scanln(&Name)
  // printing final name
	fmt.Println("Name confirmed : ", Name)
}
