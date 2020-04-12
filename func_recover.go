// Recover function is used to handle panic. 
// Only useful in deferred functions.
// Current function will not attempt to continue, but higher functions in call stack will.
// Recover function always called inside the deferred function not in the normal function. 
// If we call the Recover function as normal func then it will be not stop panicking.
// In deferred function, it will be perform at the end of program execution. Because features of defer func. 
// To take over the panic error results(error msg,goroutine traces & exit with a non zero status), we used recover function.

package main

import "fmt"

// func main(){
// 	anonymous func
// 				defer func(){
// 					str := recover()
// 					fmt.Println(str)
// 				}()
// 		panic("PANIC")
// }
func point() {
	if a := recover(); a != nil{
			fmt.Println("Recover function:- ",a)
	}
}
		// func 
func book(lang *string, name *string){
	// called as a normal function
	// point()
	// called with deferred func
	defer point()
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
	fmt.Println("Deferred function defined with normal function.")
}
// main function
func main(){
		// bookName := "Rapid Fire"
		langName := "English"
		// book(&langName,&bookName)
		book(&langName,nil)
		fmt.Println("Main function.")

}