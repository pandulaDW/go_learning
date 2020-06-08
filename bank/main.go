package main

// package will be searched in GOPATH/src
// src file name doesn't have to match the package name
// all files in a directory should have the same package
// and package can be spread through multiple files

// when there are name clashes use aliases

import (
	"fmt"
	"mypackages/numbers"
	s "mypackages/strings"
	"strings"
)

func main() {
	var num int = 40
	fmt.Printf("%d is even: %t \n", num, numbers.Even(num))
	fmt.Printf("%d is a prime: %t \n", num, numbers.IsPrime(num))
	fmt.Println(strings.Repeat("#", 20))
	str := []string{"Hope", "this", "will", "work"}
	fmt.Println(s.ConcatWithSpace(str))
}
