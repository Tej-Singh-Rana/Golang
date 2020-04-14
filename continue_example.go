// continue statement, whenever it's encountered in a loop, execution get interrupted and "continue" passes
// immediately back to the top of the loop. another execution, iteration, and any post-statement in a loop
// also executed as a part of continue statement.

package main
	
import (
		"fmt"
		"time"
)

func main(){
// pre-statement --> used once(timer := 10), post-statement --> used end of the loop over(timer--).
			for timer := 10; timer >= 0; timer-- {
				if timer%2 == 0 {
					continue  // interrupt the loop flow, right back to the top. Not print the even number.
				}
				fmt.Println(timer)
				time.Sleep(1 * time.Second)
				
			}
}