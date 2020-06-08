package main

import (
	"fmt"
)

// can be asserted to any value
type emptyInterface interface{}

type person struct {
	info interface{} // can store anything here
}

// if you don't know the type
func printSpecial(v interface{}) {
	switch v.(type) {
	case int:
		fmt.Println("This is an int")
	case string:
		fmt.Println("This is a string")
	case []string:
		for _, value := range v.([]string) {
			fmt.Println(value)
		}
	default:
		fmt.Printf("%T %v \n", v, v)
	}
}

func main() {
	var empty interface{}
	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty)

	empty = []int{12, 43, 54}
	fmt.Println(empty)
	// fmt.Println(len(empty)) cannot use empty interfaces directly for operations

	// to do this, have to do assertion
	// to access to an interface value's underlying concrete value
	fmt.Println(len(empty.([]int)))

	you := person{}
	you.info = "Your name"
	you.info = 40
	you.info = []float64{2.4, 2.3}
	fmt.Println(you.info)

	var v interface{}
	v = []string{"hi", "hiiii", "heyyyyyy"}
	printSpecial(v)
}
