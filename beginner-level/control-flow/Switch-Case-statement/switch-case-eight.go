package main

import "fmt"

func main() {
	// Main code block
	logStart := "success"
	// logStart = "fail"
	
	switch logStart {

	case logStart:
		if logStart == "error" || logStart == "fail" {
			fmt.Println("Found an error!!")
		} else if logStart == "success" || logStart == "pass" {
			fmt.Println("Process passed!!")
		} else {
			fmt.Println("Uknown bug.")
		}
		// It will pass to the next case even if case 1 is successfully executed.
		fallthrough

	case logStart:
		fmt.Println("Printing the value of logStart :", logStart)

	}

}

/*
_Output_:-
Process passed!!
Printing the value of logStart : success


*/
