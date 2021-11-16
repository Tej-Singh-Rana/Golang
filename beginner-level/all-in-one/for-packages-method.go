package main

import (
	f "fmt"
	"os/exec"
	"runtime"
	"time"
)

func track() {

	for i := 0; i < 1000; i++ {

		f.Println("Running docker command in a loop")

		call, err := exec.Command("docker ps").Output()

		if err != nil {
			f.Printf("%s", err)
		}

		output := string(call[:])
		f.Println(output)
		time.Sleep(5 * time.Second)

	}
}

func main() {

	// Main code block
	if runtime.GOOS == "Windows" {
		f.Println("It cannot be run on windows machine.")
	} else {
		track()
	}

}

