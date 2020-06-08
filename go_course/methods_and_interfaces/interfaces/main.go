package main

import (
	"fmt"
	"math"
)

// If a type implement all the method in an interface,
// it will be implementing the interface implicitly, no need to
// explicitly implement interface like in java

// Interface Type
type shape interface {
	area() float64
	perimeter() float64
}

// Rectangle type
type rectangle struct {
	width, height float64
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}

// Circle type
type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

// additional method for c
func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

// we can do polymorphism, otherwise we would have to write 2 functions
// for circle and rectangle
func printShape(s shape) {
	fmt.Println("Shape: ", s)
	fmt.Println("Area: ", s.area())
	fmt.Println("Perimeter: ", s.perimeter())
}

func main() {
	// s can take any type which implements its type
	// in go variable types are known at compilation type,
	// however interfaces can be dynamically typed and can be changed at runtime
	var s shape // 0 value of interface type is nil
	fmt.Printf("%T\n", s)

	ball := circle{radius: 2.5}
	s = ball

	printShape(s)
	fmt.Printf("Type of s: %T\n", s)

	room := rectangle{width: 2, height: 3}
	s = room
	fmt.Printf("Type of s: %T\n", s)

	// type assertions
	//  A type assertion provides access to an interface value's underlying concrete value.
	var s1 shape = circle{radius: 2.5}
	fmt.Printf("%T \n", s1) // circle

	// s1.volume() // cannot do this, only interface's methods can be called
	// have to do type assertion for this
	s1.(circle).volume()

	ball, ok := s1.(circle) // if type assertion works, ok will be true
	if ok == true {
		fmt.Printf("Ball Volume: %v\n", ball.volume())
	}

	s1 = rectangle{width: 3.4, height: 2.2}
	// type switches
	switch value := s1.(type) {
	case circle:
		fmt.Printf("%#v has circle type", value)
	case rectangle:
		fmt.Printf("%#v has rectangle type", value)
	}
}
