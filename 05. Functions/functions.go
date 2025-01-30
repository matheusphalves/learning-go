package main

import "fmt"

func main() {

	name := "Alfred"
	ages := []int{10, 20, 30}
	fmt.Println(length(name))
	fmt.Println(sum(1, 2))
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(ages...))
}
// Pass by value
func length(s string) int {
	return len(s)
}

func sum2(x int, y int) int {
	return x + y
}

func sum(x ...int) (int, int) {
	sum := 0 
	for _, value := range x {
		sum+=value
	}

	return sum, len(x)
}
