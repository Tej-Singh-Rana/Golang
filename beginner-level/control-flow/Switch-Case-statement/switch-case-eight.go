package main

import "fmt"

func main() {
	// Main code block
	logStart := "success"
	// logStart = "fail"
	
	switch logStart {
        // case 1
	case logStart:
		if logStart == "error" || logStart == "fail" {
			fmt.Println("Found an error!!")
		} else if logStart == "success" || logStart == "pass" {
			fmt.Println("Process passed!!")
		} else {
			fmt.Println("Uknown bug.")
		}
		// It will pass to the next case (i.e. case 2) even if case 1 is successfully executed.
		fallthrough
        // case 2
	case logStart:
		fmt.Println("Printing the value of logStart :", logStart)

	}

}

/*
_Output_:-
Process passed!!
Printing the value of logStart : success


*/
