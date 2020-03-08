package main
import "fmt"
			
func main() {
  //fmt.Println(4+4)
	var v1 uint= 2								//mismatched types int and uint, uint and int types must be same otherwise will occur conflicts.
	var v2 uint = 10						// 3 concepts of variables in Go ---> 1. Name of the variable, 2. Type of the variable and 3. Data
	var v3 int = -10
	var v4 int = -5          // uint ---> constant -10 overflows uint
	var v = v1 + v2
	var x = v3 + v4						// uint supported all positive values(+,0) and int supported all integers values(+,-).
	fmt.Print(v,"\n")					//"Print" func print the value in same line and cannot move into the next line.
	fmt.Println(x)            //"Println" func print the value and move into the next line.
	fmt.Println("Total value of v1 and v2 = ", v)
	fmt.Print("Total value of v3 and v4 = ", x)
}