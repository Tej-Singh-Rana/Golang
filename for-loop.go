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
// create a for loop by using keyword "for", providing a conditional expression which is either true or false.
// if condition is false it will jump into the next line of our program
// if condition is true then it will run the statements inside of the block{}.
/* forever run
package main

func main()  {
	for {

	}
}*/