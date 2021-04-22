# For Loop

#### Format

```go
for Initialization, condition, count(++,--) {
   // code block
   
}
```

```go
// "while" loop

for condition {
// Statement

}
```

```go
// Infinite loop

for {
// Code block

}
```

#### Sample Code

```go
package main

func main() {

	for i := 0; i < 5; i++ {
		// "\n" for new line.
		print(i, "\n")
	}
}
```

```go
package main

func main() {
	x := 2
	for x == 2 {
	         // "while" loop
		// "\n" for new line.
		print(x, "\n")
	}
}
```

```go
package main

func main() {
	w := 10
	// Infinite loop
	for {
		// "\n" for new line.
		print(w, "\n")
	}
}
```
