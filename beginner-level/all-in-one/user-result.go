package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

var (
	name  string
	score string
)

func main() {
	// frontend
	// taking input from user
	user := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name: ")
	name, _ = user.ReadString('\n')

	// taking score from user
	user = bufio.NewReader(os.Stdin)
	fmt.Println("Enter your total score: ")
	score, _ = user.ReadString('\n')
	cal, _ := strconv.ParseFloat(strings.TrimSpace(score), 64)
	var integer int = int(cal)
	// backend
	if integer <= 50 {
		fmt.Printf("Hello, %v\nYour score is %d.\n", name, integer)
		fmt.Println("Your score is below or equal to 50 so you couldn't qualify this time.")
	} else if integer < 85 {
		fmt.Printf("Hello, %v. Your score is %d.\n", name, integer)
		fmt.Println("Congratulations, You are shortlisted.")
	} else if integer >= 85 {
		fmt.Printf("Hello, %v. Your score is %d.\n", name, integer)
		fmt.Println("Congratulations, You are in final list.")
	} else {
		fmt.Println("Enter correct score value.")
	}

	fmt.Println("")
	fmt.Printf("Time: %v and OS Type: %v.", time.Now().Format(time.Kitchen), runtime.GOOS)

}

/*
_Output_:-
Enter your name:
john
Enter your total score:
87
Hello, john
. Your score is 87.
Congratulations, You are in final list.

Time: 3:04PM and OS Type: Linux.
*/
