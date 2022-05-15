// creating our own error
package main

import (
		"errors"
		"fmt"
)

func	main(){

	err := errors.New("Invalid Page 404")
	fmt.Println(err)
}