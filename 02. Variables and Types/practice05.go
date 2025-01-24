package main

import "fmt"

type my_type int

var x my_type
var y int

func main() {
	fmt.Printf("%v -> %T\n", x, x)
	x = 42
	fmt.Printf("%v -> %T\n", x, x)
	y = int(x)
	fmt.Printf("%v -> %T", y, y)

}
