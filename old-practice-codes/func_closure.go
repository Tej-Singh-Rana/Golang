// possible to create functions inside functions.
// Go supports anonymous function.
// An anonymous function is a function which doesn’t contain any name.
//  It is useful when you want to create an inline function.
// In Go language, an anonymous function can form a closure. 
// An anonymous function is also known as `function literal`.
/* func(parameter_list) (return_type){
-- code --
-- use return statement if return type is given
-- if return type is not given then do not use
-- return statement.
return
}()
*/
// local function can access local variables inside func.

package main

import "fmt"
// func inside func, returns func.
// defined inline func without name.
func val() func() int{
			i := 0
			return func() int{
				i++
				return i
			}
}
func main(){
		// Assigning an anonymous   
    // function to a variable   
		add := func(x, y int)int{
			return x+y
		}
		fmt.Println(add(15,20))
		// assigning in a variable
		anony := val()
		fmt.Println(anony())	
		fmt.Println(anony())
		// val()	empty result
}	