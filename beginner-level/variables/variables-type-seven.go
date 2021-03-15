package main

import (
	"fmt"
	"strconv"
)

var core int

func main() {
    // Specify type and perform a conversion
    
	core = 12
	core1 := strconv.Itoa(core)
	fmt.Printf("core1 value = %s and core1 type = %T", core1, core1)

}
