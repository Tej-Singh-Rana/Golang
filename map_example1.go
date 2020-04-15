// declaration and initialization of the map using the composite literal form.
// when passed it into the function they get passed automatically by reference rather then being copied.
// slice || pointers || maps are references.


package main

import "fmt"

func main() {
		for index, value := range map[string]string{
							"one":      "declaration",
							"two": 			"and",
							"three":		"initialization",	
							"four":			"of the map",
							"five": 		"using the composite",
							"six":			"literal",
							"seven":		"form",
}{
					fmt.Println("Index is :",index,"Value is :",value)
					
}

}