package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"runtime"
	"strings"
	"sync"
)

func checkAndSaveBody(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		fmt.Printf("%s is DOWN!!\n", url)
	} else {
		defer resp.Body.Close()
		fmt.Printf("%s -> Status Code: %d \n", url, resp.StatusCode)

		if resp.StatusCode == 200 {
			bodyBytes, err := ioutil.ReadAll(resp.Body)
			file := strings.Split(url, "//")[1] // http://www.google.com
			file += ".html"

			fmt.Printf("Writing response body to %s \n", file)
			err = ioutil.WriteFile(file, bodyBytes, 0644) // a very quick way to write to a file

			if err != nil {
				log.Fatal(err)
			}
		}
	}

	wg.Done()
}

func main() {
	urls := []string{"https://golang.org", "https://www.google.com", "https://www.medium.com"}

	var wg sync.WaitGroup
	wg.Add(len(urls))

	for _, url := range urls {
		go checkAndSaveBody(url, &wg) // creating go routines
		fmt.Println(strings.Repeat("#", 20))
	}

	fmt.Println("No. of Goroutines: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Program is done")
}
