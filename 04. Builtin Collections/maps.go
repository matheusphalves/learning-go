package main

import "fmt"

func main() {

	friends := map[string]int{
		"alfed": 123,
		"mark":  456,
	}

	fmt.Println(friends)
	fmt.Println(friends["mark"])

	friends["gopher"] = 1233

	fmt.Println(friends)
	fmt.Println(friends["gopher"])

	if friend, exists := friends["alan"]; !exists {
		fmt.Println("alan don't exists")
	} else {
		fmt.Println(friend)
	}

	friend, exists := friends["alan"]
	fmt.Println(friend, exists)

	anotherMap := map[int]string{
		123:  "cool",
		456:  "awesome",
		1234: "party time!",
		676:  "someone",
	}

	for key, value := range anotherMap {
		fmt.Println(key, value)
	}

	delete(anotherMap, 456)
	delete(friends, "mark")

	fmt.Println(anotherMap)
	fmt.Println(friends)

}
