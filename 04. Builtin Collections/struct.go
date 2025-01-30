package main

import "fmt"

type Product struct {
	name  string
	price float32
	isNew bool
}

type Customer struct {
	name     string
	age      int
	products []Product
}

func main() {

	productA := Product{
		name:  "Headset",
		price: 50.5,
		isNew: true,
	}

	productB := Product{"Keyboard", 20.0, false}

	fmt.Println(productA)
	fmt.Println(productA.name)
	fmt.Println(productB)

	customerA := Customer{
		name:     "Alfred",
		age:      20,
		products: []Product{productA, productB},
	}

	fmt.Println(customerA)
	fmt.Println(customerA.products[0])

	// Annonymous struct

	person := struct {
		name string
		age  int
	}{
		name: "Alfred",
		age:  28,
	}

	fmt.Println(person)

}
