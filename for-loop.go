package main

import "fmt"

func main()  {
		sum := 0
		//Initialization(i:=0) ; Condition(i<10); Update(i++)
	for i:= 0; i<10; i++{
		sum +=i
		fmt.Println(i)
		fmt.Println(sum)
	}
	fmt.Println(sum)
}

/* forever run
package main

func main()  {
	for {

	}
}*/