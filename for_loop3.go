package main

import "fmt"

func main()  {
		i := 1
		// condition expression either true or false
		for i <= 10 {
			fmt.Println(i)
			i += 1	// extremely important this section without this "for" loop will print infinite time.
							// i <= 10 always evaluate to true and loop will never stop.
		}
//  variable initialization; condition to check each time; increment the variable (increment form --> x++, decrement form --> x--) 
	  for x := 1; x <= 10; x++ {
				fmt.Println(x)
		}
		for y := 1; y >= -10 ; y-- {
				fmt.Println(y)
		}
}