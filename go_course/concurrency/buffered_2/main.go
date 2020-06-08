package main

import (
	"fmt"
	"runtime"
	"time"
)

// won't lock the goroutine until buffer is full

func main() {
	c := make(chan int, 3)

	go func(c chan int) {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Func goroutine #%d starts sending data into the channel\n", i)
			c <- i
			fmt.Printf("Func goroutine #%d after sending data into the channel\n", i)
		}
		close(c)
	}(c)

	fmt.Println("main goroutine sleeps for 2 seconds")
	time.Sleep(time.Second * 2)

	fmt.Printf("No of goroutines running: %d\n", runtime.NumGoroutine())

	fmt.Println("main goroutine starts receiving data")
	for v := range c {
		fmt.Println("main goroutine received data: ", v)
	}

	fmt.Println(<-c) // when trying to receive a value from a closed channel, it wil give zero value

	c <- 10 // sending data to a closed channel will make the program panic
}
