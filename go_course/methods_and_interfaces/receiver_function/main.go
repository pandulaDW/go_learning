package main

import (
	"fmt"
	"math"
	"time"
)

// methods belong to a type and functions belong to a package
type names []string

// n is called receiver argument
func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

// Vertex of struct type
type Vertex struct {
	X, Y float64
}

// Abs method of type Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T %v\n", day, day) // type = time.Duration

	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds)
	fmt.Printf("Seconds in a day: %v \n", seconds)

	friends := names{"Dan", "Marry"}

	// Both works
	friends.print()
	names.print(friends)

	var n int64 = 9523123
	fmt.Println(n)
	fmt.Println(time.Duration(n))

	v := Vertex{3., 4.}
	fmt.Println(v.Abs())
}
