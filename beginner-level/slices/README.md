# Slices 

#### Sample Code

```go
func main() {
   // Type = int, string, bool, float
   // Add Elements = 10, "world", 12.20
   foo := []Type{ Add Elements }  
   
}
```
> Note:- All elements should be same type. 

```go
package main

import (
	"fmt"
)

func main() {
	// Slice
	// []type{ Length is not fixed }

	exampleCode := []string{"sample", "code", 2021}

	fmt.Println(exampleCode)
}

__Output__:- cannot use 2021 (type untyped int) as type string in slice literal
```
