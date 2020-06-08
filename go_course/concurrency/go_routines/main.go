package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func f1(wg *sync.WaitGroup) {
	fmt.Println("f1 (goroutine) execution started")
	for i := 0; i < 3; i++ {
		fmt.Println("f1, i = ", i)
		time.Sleep(time.Second) // simulate expensive task
	}
	fmt.Println("f1 (goroutine) execution finished")
	wg.Done() // same as (*wg).Done()
}

func f2() {
	fmt.Println("f2 (goroutine) execution started")
	for i := 5; i < 8; i++ {
		fmt.Println("f2, i = ", i)
	}
	fmt.Println("f2 (goroutine) execution finished")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1) // no of goroutines to wait for

	go f1(&wg) // passing a pointer, since we are mutating this wg variable
	fmt.Println("No. of Goroutines after go f1(): ", runtime.NumGoroutine())

	f2()

	// BLock the main executing thread, until wg is completed
	wg.Wait()
	fmt.Println("main executing stopped")
}
