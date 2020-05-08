package main


import "fmt"

var (
		x = 12  int,
		y = 15  int,
		z = 20  int,

)

type Number struct {
				count int
}

func Strange(h Number) func() int{
			add := x + y + z
		return add
}

func main() {

		c := Strange()
		fmt.Println(c)
		fmt.Println(c.y)

}