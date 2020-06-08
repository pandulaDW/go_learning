package main

import "fmt"

func f1(n int, ch chan int) {
	ch <- n
}

func main() {
	// by directional channel
	c := make(chan int)

	// only for receiving channel
	c1 := make(<-chan string)

	// only for sending
	c2 := make(chan<- string)

	fmt.Printf("%T, %T, %T \n", c, c1, c2)

	go f1(10, c)

	n := <-c
	fmt.Println("Value received: ", n)
	fmt.Println("Exiting main...")
}
