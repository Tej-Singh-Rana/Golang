// map function

package main

import "fmt"

func main(){
// initializing variable
	var x = map[int]string{
		1: "Hello",
		2: "Good",
		3: "Mid-Night",
		4: "Map",
		5: "Function",
		6: "Initializing",
	}
	fmt.Println(x)
	for id, comment := range x {
		fmt.Println("ID",id,"comment",comment)
	}
}