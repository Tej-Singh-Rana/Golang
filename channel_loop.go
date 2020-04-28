package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	wg.Add(2)
	go func(ch <-chan int, wg *sync.WaitGroup) {
		// receiving message
		// for i := 0; i < 10; i++ {
		// 	fmt.Println(<-ch)
		// }
		for msg := range ch {
			fmt.Println(msg)
		}
		wg.Done()
	}(ch, wg)

	go func(ch chan<- int, wg *sync.WaitGroup) {
		// sending message
		for i := 1; i <= 10; i++ {
			ch <- i

		}
		// after complete the loop, closing the sender channel.
		// because it's sending a message and receiver side listening new messages but not genrating so occur a deadlock situation.
		// after close function go routine will detect there is not message to route anymore.
		close(ch)
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
