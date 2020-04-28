/* Channels send and receive messages through go channels, after sending a messages to go channel.
When from receiver side, get a messages but cannot determine that who sent a messages, rest functionality
has performed by functions of receiver.

*/

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	wg := &sync.WaitGroup{}
	ch := make(chan string, 1)

	wg.Add(2)
	// anonymous function
	// receive-only channel
	go func(ch <-chan string, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		// ch <- "hello"   // invalid operation for this block
		wg.Done()

	}(ch, wg)
	//  send-only channel
	go func(ch chan<- string, wg *sync.WaitGroup) {
		ch <- "Connecting to server...."
		time.Sleep(2 * time.Millisecond)
		ch <- "Server successfully started."

		wg.Done()
	}(ch, wg)

	wg.Wait()
}
