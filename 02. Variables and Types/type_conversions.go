package main

import "fmt"

type hotdog int

var b hotdog = 10

func main() {
	x := 10
	// fmt.Printf("%T\n", b)
	// fmt.Printf("%T", x)
	fmt.Printf("%v\n", x)
	fmt.Printf("%v\n", b)

	x = int(b)
}
