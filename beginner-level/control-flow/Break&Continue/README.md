# Break & Continue

### Break 

#### Format

```go
func main() {

	for <initialization>;<condition> {
		if <condition> {
                        // code block
			break   // --> Exit from the loop
		} else {
			// code block
		}

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

	for i := 10; i <= 20; i++ {
		if i == 15 {
			fmt.Println(i)
			break   // --> Exit from the loop.
		} else {
			fmt.Println(i)
		}

	}

}
```

### Continue

#### Format




#### Sample Code






