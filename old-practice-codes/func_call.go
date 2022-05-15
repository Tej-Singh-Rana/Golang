package main

import "fmt"

func sin(number int) {
	fmt.Println("First one calling...", number)

}

func cos(name string) {
	fmt.Println("Second one calling...", name)

}

func main() {
	sin(1)
	cos("start")
}
