package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		// receiving message
		// fmt.Println(<-ch)
		// declare variable to get the status of channel. Sender channel is close or open.
		// _, ok := <-ch
		// fmt.Println("Channel status :", ok)
		// with control flow, If statement
		// In this case, it prints only when channel is open.
		if msg, ok := <-ch; ok {
			fmt.Println(msg, ok)
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		// sending message
		ch <- 12345789656
		// close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
