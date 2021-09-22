// 

package main

import (
	"fmt"
	"time"
)

func main() {
	// Main code block
	for i := 0; i < 100000; i++ {
		fmt.Println("Testing ", i)
		time.Sleep(2 * time.Second)
		fmt.Println("Time is now", time.Now().Format("01-02-2006 15:04:05"))
	}
}

/*
_Output_:-
Testing  0
Time is now 09-22-2021 23:37:40
Testing  1
Time is now 09-22-2021 23:37:42
Testing  2
....continue

*/
