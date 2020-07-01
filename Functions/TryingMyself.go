package main

import "fmt"

func main() {
	x := bar()
	fmt.Printf("%T\n", x)

	// x()Run func
	i := x()
	fmt.Println(i)
}

func bar() func() int {
	return func() int {
		return 451
	}
}
