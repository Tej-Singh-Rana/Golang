// In generally, it's occur when program cannot continue at all.
// It's use for unrecoverable events.
// If nothing handles panic, program will exit.
// When function will stop executing , deferred functions will still works.
/* Running this program will cause it to panic, print an error message and goroutine traces, and 
exit with a non-zero status.
*/
package main

import (
			"fmt"
)
// func 
func book(lang *string, name *string){
		// when lang section is nil it will be panic
		if lang == nil{
				panic("Error: Language section cannot be empty!!")
		}
		// when name section is nil it will be panic
		if name == nil {
				panic("Error: Name section cannot be empty!!")
		}
		// when value is not nil then it will be print
		fmt.Printf("Book language: %s\n Book name: %s\n ", *lang, *name)
}
// main function
func main(){
			// bookName := "Rapid Fire"
			langName := "English"
			// book(&langName,&bookName)
			book(&langName,nil)

}