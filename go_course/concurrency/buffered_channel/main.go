package main

import "fmt"

func squares(c chan int) {
	for i := 0; i < 3; i++ {
		num := <-c
		fmt.Println(num * num)
	}
}

func main() {
	fmt.Println("main() started")
	c := make(chan int, 3) // this is a buffered channel

	go squares(c)

	//  goroutine is not blocked until after buffer is full
	c <- 1
	c <- 2
	c <- 3
	c <- 4 // blocks here

	// Length of a channel is the number of values queued (unread) in channel
	// Capacity of a channel is the buffer size
	fmt.Printf("Length and capacity of channel c is %v, %v\n", len(c), cap(c))
	fmt.Println("main() stopped")
}
