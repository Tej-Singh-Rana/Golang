// switch case examples
// comparing switch and case types
// switch case runs top to down and run each cases end of the last curly brackets.
// you cannot assign one case multiple times, otherwise gets the duplicate case error.
// implicit breaks
// requires explicit fallthrough

package main

import "fmt"

const (
				lockDown = "April"

)

func main(){
			switch lockDown {
			case "April": 
				fmt.Println("Lockdown-1 completed at 14 April.")
				fallthrough
			case "May": 
				fmt.Println("Lockdown-2 started from 14th April to 3rd May.")
				fallthrough
			case "June": 
				fmt.Println("Lockdown-3 pending.......")
			
			case "July":
				fmt.Println("Lockdown-4 pending........")
			
			case "August":
				fmt.Println("You wanna take this whole year?? Right now close this cases.")
			
			default:
					fmt.Println("Hardest time ever.")
			
			}
}