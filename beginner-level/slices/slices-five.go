package main

import (
	"fmt"
)

func main() {
	// Main code block
	group := []string{"class1", "class2", "class3", "class4"}

	for _, value := range group {

		fmt.Println("List of classes - ", value)

	}

	// If we are trying to filter only "class1"

	fmt.Println(group[0])
	
	// If filtering all items after "class2"
	
	fmt.Println(group[1:])

}

/*
_Output_:-
List of classes -  class1
List of classes -  class2
List of classes -  class3
List of classes -  class4
class1
[class2 class3 class4]

*/
