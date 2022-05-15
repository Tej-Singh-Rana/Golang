// append method used

package main

import "fmt"

func main()  {
	x := []string{"tyson","car","image","color"}
	y := append(x,"house","business","tower")
	z := make([]string,7)
	copy(z,y)
	fmt.Println(x,y,z)
}