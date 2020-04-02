/* Variables assigning format :--
var name string = "Hello World"
*/
package main

import "fmt"

func main()  {
		// keyword var is for assign a variables.
		var name = "Go"
		fmt.Println("name is assigned the string Go ",name)
//  name := name + " programming language"  // It throws the error no new variables on left side of :=, already defined. 
		name = name + " programming language"
		fmt.Println("Name of the programming language : ", name)
//  In GO, compiler is able to inference var types. var x = "compiler" or x := "compiler" or x := 5
		language_name := "Go Language"
		fmt.Println("Name of language : ", language_name)
		languageName := "Go"     //  lower camel case "languageName"
		fmt.Println(languageName)

}

