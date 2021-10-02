// https://play.golang.org/p/nL7SYcgGam0

package main

import (
	"fmt"
	"strings"
)

func main() {
	// Main code block

	var Ground []string = []string{"latency", "axis", "longitude", "latitude"}

	fmt.Println(Ground)

	// strings package, delimiter used as comma

	fmt.Println(strings.Join(Ground, ","))

}

/*
_Output_:-
[latency axis longitude latitude]
latency,axis,longitude,latitude

*/
