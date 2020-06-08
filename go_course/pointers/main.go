package main

import (
	"fmt"
	"strings"
)

// go always pass by value. never by reference
// slice, map pointers are also copied (only those two, not struct)
// to modify a variable outside function scope, you need to pass its pointer

func changeValues(quantity int, price float64, name string, sold bool) {
	quantity = 3
	price = 500.4
	name = "Mobile Phone"
	sold = false
}

func changeValuesByPointer(quantity *int, price *float64, name *string, sold *bool) {
	*quantity = 3
	*price = 500.4
	*name = "Mobile Phone"
	*sold = false
}

// Product is a struct type of products
type Product struct {
	price       float64
	productName string
}

// doesn't modify the struct
func changeProduct(p Product) {
	p.price = 300
	p.productName = "Bicycle"
}

func changeProductByPointer(p *Product) {
	(*p).price = 300 // equivalent to p.price = 300
	p.productName = "Bicycle"
}

func changeSlice(s []int) {
	for i := range s {
		s[i]++
	}
}

func changeMap(m map[string]int) {
	m["a"] = 10
	m["b"] = 20
	m["c"] = 30
}

func main() {
	quantity, price, name, sold := 5, 300.4, "Laptop", true
	fmt.Println("Before calling changevalues()", quantity, price, name, sold)
	changeValues(quantity, price, name, sold)
	fmt.Println("after calling changevalues()", quantity, price, name, sold)

	strings.Repeat("-", 50)

	fmt.Println("Before calling changevalues()", quantity, price, name, sold)
	changeValuesByPointer(&quantity, &price, &name, &sold)
	fmt.Println("after calling changevalues()", quantity, price, name, sold)

	gift := Product{price: 100, productName: "Watch"}
	changeProduct(gift)
	fmt.Println(gift)

	changeProductByPointer(&gift)
	fmt.Println(gift)

	prices := []int{1, 2, 3}
	changeSlice(prices)
	fmt.Println(prices)

	myMap := make(map[string]int)
	changeMap(myMap)
	fmt.Println(myMap)
}
