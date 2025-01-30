package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (person Person) helloMethod(greetings string) {
	fmt.Println(person.name, "says: ", greetings)
}

type Dentist struct {
	person  Person
	address string
}

type Engineer struct {
	person         Person
	registerNumber int
}

func (x Dentist) greetings() {
	fmt.Println("Dentist, ", x.address)
}

func (x Engineer) greetings() {
	fmt.Println("Engineer, ", x.registerNumber)
}

type Human interface {
	greetings()
}

func humanGreetings(human Human) {

	switch human.(type) {
	case Dentist:
		fmt.Println(human.(Dentist).address)

	case Engineer:
		fmt.Println(human.(Engineer).registerNumber)
	}

	human.greetings()
}

func main() {

	dentist := Dentist{
		person:  Person{"AlfredDentist", 25},
		address: "My Address",
	}

	engineer := Engineer{
		person:         Person{"EngineerName", 30},
		registerNumber: 12342,
	}

	humanGreetings(dentist)
	humanGreetings(engineer)

	dentist.greetings()
	engineer.greetings()

}
