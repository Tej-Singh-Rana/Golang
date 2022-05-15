package main

import "fmt"

func main() {
	// channels are blocking constructs.
	ch := make(chan int)

	// case 1
	// trying to receive a message from a channel, but haven't generated yet.
	// generate the deadlock condition.
	fmt.Println(<-ch)
	/* we receive a message from a channel but cannot pass the message in above channel.
	we cant get there, until a message comes in.
	*/
	ch <- 500
	// if we do in reverse order
	// case 2
	// It cannot receive the message into the channel until there's a receiver.
	// pushing message into the channel but cannot pass it.
	ch <- 500
	// it will not listen until it's have receiver.
	// senders & receiver
	fmt.Println(<-ch)

}
