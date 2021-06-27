package main

import (
	"fmt"
	"time"
)

func main() {
	var (
		firstCount = 0
		lastCount  = 10
		rollout    = "fast"
	)

	for firstCount <= 10 {
		if firstCount == lastCount {
			fmt.Println("This one is last roll out.\nEnd")
			break   // We don't need 11th rollout. So when it's match with the lastCount value. It will break the loop. 
		}
		fmt.Println("Counting...")
		time.Sleep(2)
		firstCount++
		fmt.Println(firstCount)
		fmt.Println("Rollout Status", rollout)

	}

}

/*
Output :-
Counting...
1
Rollout Status fast
Counting...
2
Rollout Status fast
Counting...
3
Rollout Status fast
Counting...
4
Rollout Status fast
Counting...
5
Rollout Status fast
Counting...
6
Rollout Status fast
Counting...
7
Rollout Status fast
Counting...
8
Rollout Status fast
Counting...
9
Rollout Status fast
Counting...
10
Rollout Status fast
This one is last roll out.
End
*/
