package main

import (
	"fmt"
	"sync"
	"time"
)

// any code that is present between a call to lock and unlock will be
// executed by one one goroutine

func main() {
	const gr = 100
	var wg sync.WaitGroup
	wg.Add(gr * 2)

	var n int = 0

	var m sync.Mutex

	for i := 0; i < gr; i++ {
		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			n++ // reading, incrementing and assigning will not be interrupted by another goroutine
			m.Unlock()
			wg.Done()
		}()

		go func() {
			time.Sleep(time.Second / 10)
			m.Lock()
			defer m.Unlock()
			n-- // reading, decrementing and assigning will not be interrupted by another goroutine
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value of n: ", n)
}
