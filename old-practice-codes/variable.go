package main

import ("fmt")

func main()  {
		  var volume1, volume2 int
			
			//var volume1 int 
			//var volume2 int
		 volume1, volume2 = 10, 15
		 var volume = volume1 + volume2
		 fmt.Println(volume)
		 fmt.Println("------------------------------------")
		 var pass = 10
		 pass = 20
		 fmt.Println(pass)
		 fmt.Println("------------------------------------")
		 value := 10       // same like var value = 10
		 value = 90					// we can change the value of that variable
		 fmt.Println(value)
		 fmt.Println("-------------------------------------")
		const num = 20       //to fix the value of that variable, cannot be changed so we used "const" keyword
	//num = 30							// assign to "num" here, it gives an error that cannot assign to num again
		fmt.Println(num)
}