//
//

package main

import "fmt"

func main(){
		language := []string{
			"Java",
			"Python",
			"Go",
			"JavaScript",
			"C#",
			"C++",
			"C",
			"Perl",
		}

		lang := []string{
			"English",
			"Hindi",
			"Spanish",
			"French",
			"Japanese",
			"British",
			"Russian",
			"Chinese",
			"Python",
			"Dutch",
		}
		// outer loop
		// pick the value one by one of var language, and matches the pattern in the inner loop.
		for _, i := range language{
			// data value
			fmt.Println(i,"--------")
			// inner loop
			for _, j := range lang{
				// inner loop rolling out till all values of var lang are not performed.
				fmt.Println(j)
				if i == j {
					fmt.Println("Python language available in both category.")
					break
				}
			}
		}
}