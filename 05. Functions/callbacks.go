package main

import "fmt"

func main() {

	sum := onlyEvenNumbers(funcSum, []int{50, 51, 52, 53, 54}...)
	fmt.Println(sum)
}

func funcSum(x ...int) int {
	sum := 0
	for _, value := range x {
		sum += value
	}
	
	return sum
}

func onlyEvenNumbers(f func(x ...int) int, y ...int) int {
	var slice []int

	for _, value := range y {
		if value % 2 == 0 {
			slice = append(slice, value)
		}
	}

	totalSum := f(slice...)
	return totalSum
}