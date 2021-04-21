package main

import (
	"fmt"
)

func main() {
	// Slice
	// []type{ Length is not fixed }

	examYear := []int{collegeYear(), 2021, 2022}

	fmt.Println("Exam year", examYear)
}

// Declare func collegeYear()
func collegeYear() int {
	return 2020
}
