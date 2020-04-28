package main

import (
	"fmt"
	"sync"
)

// buffer channel scenario
/* 	case scenario is if we have two messages is ready to send, but receiver is not ready then it will generate deadlock condition.
 		so we have to use internal buffer in this channel.
		We have to assign a parameter in channel. It will receive a message but it will not pass it.
		second message will not received back, and at least application is not blocking the process.
*/
func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 500
		ch <- 800
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
