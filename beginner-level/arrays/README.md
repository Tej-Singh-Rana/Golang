# Arrays [length]type{items}

- In this section, we will learn about an __`array` in `Go` language__.

### Sample Code

- Method - 1
```
package main

func main() {
	// Main Code block
	// var name [length]type
	// Example:- var rank [100]int

	var Name [Length]Type
        // In an array, Index value starts from zero.
        
         Name[index-0] = value-0
         Name[index-1] = value-1
         ....
         ....
         Name[index-n] = value-n         
}
```

- Method - 2
```
package main


func main() {
	// Main code block
	// In an array, index value starts from zero.
	
	Name := [Length]Type{value-0, value-1, value-2, value-3, ..., value-N}
	
	OR
	
	var Name [Length]Type{value-0, value-1, value-2 ...., value-N}
	var Name = [Length]Type{value-0, value-1, value-2 ...., value-N}
}
```



