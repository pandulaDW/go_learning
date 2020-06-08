package main

// import statement is file scoped, meaning that each file should import
// constants, variables and functions are packaged scoped

import "fmt"

func main() {
	fmt.Println("Scope means name visibility")
	sayHello("Pandula")

	tf := toFahrenheit(boilingPoint)
	fmt.Println("Water boling point in F: ", tf)
}
