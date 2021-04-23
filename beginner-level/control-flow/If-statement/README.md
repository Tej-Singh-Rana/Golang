# If/Else Statement

#### Format

```go
func main() {
 
   // Format of "if/else" statement
   if condition {
       // statement
   } else {
      // statement
    }
}
```
> If-Else-If Statement
```go
func main() {
 
   // Format of "if/else" statement
   if initialization; condition {
     // statement   
   } else if {
      // statement  
    } else {
      // statement
    }
}
```

#### Sample Code

```go
package main

import (
	"fmt"
)

func main() {

	// Sample code of "if/else" statement

	if x := 12; x == 10 {
		fmt.Println(x)
	} else {
		fmt.Println("Different value.")
	}
	// x value available within a block,
	// fmt.Println(x)
	// Error:- undefined: x
}
```
> Note:- Now var "x" value available within a main function. 

```go
package main

import (
	"fmt"
)

func main() {

	// Sample code of "if/else" statement
	x := 12
	if x == 10 {
		fmt.Println(x)
	} else {
		fmt.Println("Different value.")
	}

	// x value available now
	fmt.Println(x)
}
```
