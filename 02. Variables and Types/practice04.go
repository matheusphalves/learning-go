package main

import "fmt"

type my_type int

var x my_type

func main() {
	fmt.Printf("%v -> %T\n", x, x)
	x = 42
	fmt.Printf("%v -> %T", x, x)
}