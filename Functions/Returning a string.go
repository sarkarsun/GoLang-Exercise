package main

import "fmt"

func main() {
	fmt.Println("---Func with return---")
	s1 := roo()
	fmt.Println(s1)

}
func roo() string {
	s := "My String is long"
	return s
}
