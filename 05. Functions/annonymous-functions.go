package main

import "fmt"

func main() {

	func(s string) {
		fmt.Println("Hello: ", s)
	}("Alfred!")

	sumFunc := func(x int, y int) int {
		return x + y
	}

	fmt.Println(sumFunc(1, 3))

	function := returnAFunction()

	fmt.Println(function(20))
}

func doSomething(something string) { 
	fmt.Println("Hello: ", something)
}

func returnAFunction() func(int) int {
	return func(i int) int {
		return i * 10
	}
}