package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3, 4, 5}
	sliceB := []int{6, 7, 8}

	// append(slice, elements)
	joinedSlice := append(sliceA, sliceB...)

	fmt.Println(joinedSlice)
}
