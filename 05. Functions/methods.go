package main

import "fmt"

type Person struct {
	name string
	age int
}

func (person Person) helloMethod(greetings string) {
	fmt.Println(person.name, "says: ", greetings)
}

func main() {

	alfred := Person{"Alfred", 20}

	alfred.helloMethod("hello!")

}
