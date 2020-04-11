// Recursion, which method or function calls itself continuously.
// 
//

package main

import "fmt"

func fact(n int) int{
			if n == 0 {				// base code
				return 1
			}
			return n * fact(n-1)
}

func main(){
			fmt.Println(fact(5))
}
