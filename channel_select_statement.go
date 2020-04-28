// Select statement is specifically desgined for channels.
// In select statement, no predetermined order unlike switch statement (case 1 is match then rest will not be perform).
// In select statement multiple cases can be perform at the same time.
// Select statement is specifically desgined for and have the current go routine pause until one of those channels is either ready to send a messsage or receive a message for the use case.
// If any one of the case is not execute then default(non-block) will be execute.
/* Syntax
select {
case 1:
				# code block
case 2:
			# code block
default:
			# use default case for non-blocking select
}
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch2 <- "Message 2"
		time.Sleep(10 * time.Second)
	}()
	go func() {
		ch1 <- "Message 1"

	}()

	for i := 0; i < 2; i++ {
		select {

		case r1 := <-ch1:
			fmt.Println(r1)
		case r2 := <-ch2:
			fmt.Println(r2)
		default:
			fmt.Println("No message describe.")

		}
	}
}
