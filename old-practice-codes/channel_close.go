// Closeing Channel
/*
	- Closed via the built-in close function
	- Cannot check for closed channel it's closed or not.
	- After closed channel, sending a new message triggers a panic(error)
	- Receiving messages okay
		* If buffered, all buffered messages available
		* If unbuffered, or buffer empty, then you're going to receive the zero value for that channel.
	- Use comma okay syntax to check, after receiving message from that channel and then check by giving second return parameter to see that's true or false.
		* true --> open
		* false --> close
	- Closing a channel always has to be on the sending side of the operation.
	- Receiver side, If channel is closed and receiving messages go runtime throws zero value back.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	// anonymous function
	// receive-only channel
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		close(ch)
		// we get zero value back, in terms of string we get empty.
		fmt.Println(<-ch)
		wg.Done()

	}(ch, wg)
	//  send-only channel
	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 200

		// close(ch)

		// sending message after closing channel
		// panic occur
		// ch <- 201
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
