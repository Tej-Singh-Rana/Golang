package main

import (
	"fmt"
	"sync"
)

// unbuffer channel scenario

func main() {
	wg := &sync.WaitGroup{}
	ch := make(chan int)

	wg.Add(2)
	go func(ch chan int, wg *sync.WaitGroup) {
		fmt.Println(<-ch)
		wg.Done()
	}(ch, wg)

	go func(ch chan int, wg *sync.WaitGroup) {
		ch <- 500
		wg.Done()
	}(ch, wg)

	wg.Wait()
}
