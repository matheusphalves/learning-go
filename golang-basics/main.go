package main

import "fmt"

func main() {
	test := "my-string-content"
	mapTest := map[string]string{
		"key": "value",
	}

	mapTestAny := map[string]interface{}{
		"key": "value",
		"key2": 20,
	}

	var test2 interface{}

	test2 = 1
	printVariable(test2)

	test2 = "string"
	printVariable(test2)
	printVariable(test)
	printVariable(mapTest)
	printVariable(mapTestAny)

}

func printVariable(value interface{}) {
	fmt.Printf("%T, %s\n", value, value)

}

func privateFunction() {
	fmt.Printf("Hello from private func")
}

func PublicFunction() {
	fmt.Printf("Hello from public func")
}