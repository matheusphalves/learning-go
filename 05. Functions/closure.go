package main

import "fmt"

func main() {

	contextA := closure()
	fmt.Println(contextA())
	fmt.Println(contextA())
	fmt.Println(contextA())

	contextB := closure()

	fmt.Println(contextB())
	fmt.Println(contextB())
	fmt.Println(contextB())


}

func closure() func() int {
	x :=0 
	return func() int {
		x++
		return x
	}
}