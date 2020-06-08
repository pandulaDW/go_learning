package main

import "fmt"

// init and main both doesn't take any arguments
// init runs before main
// runs in the declared order
// init funcs can be used to initialize global variables

// declare an array
var numbers [10]int

func init() {
	fmt.Println("This is init #1")
	for i := 0; i < len(numbers); i++ {
		numbers[i] = i * 2
	}
}

func init() {
	fmt.Println("This is init #2")
}

func main() {
	fmt.Println("This is main")
	fmt.Printf("%#v\n", numbers)
}
