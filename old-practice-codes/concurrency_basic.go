// concurrency --> planned, structured
// concurrency vs parallelism, related yes but not a same thing. 
// goroutines --> goroutine --- safe communication --- goroutine

package main

import (
				"fmt"
				"time"
				"sync"
)

func main(){
// it makes the program to wait and run
	var waitGrp sync.WaitGroup
// Add(each for goroutine), main func is going to wait until both of our goroutines report back. they done, before continue & closes.
	waitGrp.Add(2)

// declare anonymous func
// go keyword -- turns into the goroutine
// holds the entire func code.
	go	func(){
				defer waitGrp.Done()
				time.Sleep(5 * time.Second)
				fmt.Println("Starting process...")
			}() // end of this parenthesis, will make the function self-executing.

			// go keyword -- turns into the goroutine
		//next it will run the func code, after success execution, func exited so we need to execute upper func code as well.
		go	func(){
				defer waitGrp.Done()

				fmt.Println("Waiting to execute...")
			}()
		// after sleep func code executes then it will exit from main func.
		// wait for that no more code has to execute.
			waitGrp.Wait()
} // func exit
