// for loop examples
// break, breakpoint	

	package main
	
	import (
			"fmt"
			"time"
	)

	func main(){
	// pre-statement --> used once(timer := 10), post-statement --> used end of the loop over(timer--).

		for timer := 10; timer >= 0; timer-- {
					if timer == 2{
						fmt.Println("Time almost done!!")
						time.Sleep(1 * time.Second)
					}
					fmt.Println(timer)
					time.Sleep(1 * time.Second)
					if timer == 0{
						 fmt.Println("Timeout!!")
						 break 
					}
				}
	}