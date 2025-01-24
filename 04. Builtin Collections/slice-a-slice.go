package main

import "fmt"

func main() {
	colors := []string{"red", "blue", "yellow", "green"}

	// Access all elements
	// slices := colors[:]

	slices := colors[0:3]

	fmt.Println(slices)
}
