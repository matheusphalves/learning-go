package main

import "fmt"

func main() {

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)

	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	slice2 := append(slice, 6)
	fmt.Println(slice2)

	fmt.Println(slice[3])
	slice[3] = 213425
	fmt.Println(slice[3])

	// slice[20] = 1
	// fmt.Println(slice[20])

	stringSlice := []string{"apple", "orange", "pineapple"}
	forEachSlice(stringSlice)

	stringSlice = append(stringSlice, "banana")
	forEachSlice(stringSlice)
}

func forEachSlice(slice []string) {

	// Using for each statement
	for index, value := range slice {
		fmt.Println(index, "->", value)
	}
}
