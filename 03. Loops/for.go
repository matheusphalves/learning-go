package main

import "fmt"

func main() {
	for hours := 0; hours < 12; hours++ {
		fmt.Println("Hour: ", hours)
		for minutes := 0; minutes < 60; minutes++ {
			fmt.Print(minutes, " ")
		}
		fmt.Println("\n")

	}

	x := 0

	for {
		if x < 10 {
			fmt.Println("To eternity...")
			x++
		} else {
			break
		}
	}
}
