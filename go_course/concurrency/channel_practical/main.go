package main

import (
	"fmt"
	"strings"
)

func factorial(n int, c chan int) {
	f := 1
	for i := 2; i < n+1; i++ {
		f *= i
	}
	c <- f
}

func squares(c chan int) {
	for i := 0; i < 9; i++ {
		c <- i * i
	}
	defer close(c)
}

func main() {
	ch := make(chan int)
	defer close(ch)

	go factorial(5, ch)

	f := <-ch // This is a blocking operation
	fmt.Println(f)

	for i := 1; i <= 20; i++ {
		go factorial(i, ch)
		f := <-ch
		fmt.Println(f)
	}

	c := make(chan int)
	go squares(c)

	// periodic block/unblock of main goroutine until channel close
	for {
		val, ok := <-c
		if ok == false {
			fmt.Printf("%v %v, The loop is broken\n", val, ok)
			break
		} else {
			fmt.Printf("%v %v\n", val, ok)
		}
	}

	///////// Anonymous function
	fmt.Println(strings.Repeat("#", 20))
	for i := 5; i < 15; i++ {
		go func(n int, c chan int) {
			f := 1
			for i := 2; i < n+1; i++ {
				f *= i
			}
			c <- f
		}(i, ch)
		fmt.Println(<-ch)
	}

	fmt.Println("main() stopped")
}
