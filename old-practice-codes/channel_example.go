// channel --> type of queue or sequence, after one func code complete it's send a signal of completion.
// Channels are a typed conduit through which you can send and receive values with the channel operator, <-.
// ch <- v   send v to channel ch,
// v := <-ch  receive from ch and assign value to v 
// data type must be same, goroutine == channel count, keyword chan, 
// var variable-name chan type or by shorthand/make func "variable-name := make(chan type)

package main

import "fmt"

// initialization of channel
func  multiple(c chan int, value int){
	// sending data
				c <- value * 10
}

func t1(ch chan int){

			fmt.Println(250 + <-ch)
			// close channel after process over
			// defer close(ch)
}

func main(){
			// define channel

			codeValue := make(chan int)
			
			// signal sends codeValue to channel
			// declare goroutine
			go multiple(codeValue, 20)

			// receiving data
			store := <-codeValue			
			fmt.Println(store)

			// declare goroutine	
			go t1(codeValue)
			// sending data
			codeValue <- 50

}
