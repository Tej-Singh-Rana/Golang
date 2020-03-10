package main

import "fmt"

func main()  {
	code := 2 
	if  code <= 5 {
			fmt.Println("This is less then or equal to 5")
	} else if code == 10 {
	 		 fmt.Println("This is more then value of code")
	} else if code > 5 {
			fmt.Println("This is greather then 5")
	} else {												//We have to write else statement after ending of }.
		 	fmt.Println("Try again for success")
}	
		fmt.Println("Value of code will print here ",code)
	}