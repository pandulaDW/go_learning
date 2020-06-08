package main

import (
	"fmt"
	"math"
)

type car struct {
	brand string
	price int
}

// normal function
func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

// method of car type
func (c car) changeCar(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

// pointer receiver method
func (c *car) changeCarWithPointer(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

type distance *int

// methods can only be created for value types and not pointer types
// func (d distance) m1() {
// }

// Vertex type example
type Vertex struct {
	X, Y float64
}

// Abs method of type Vertex
func (v Vertex) Abs() float64 {
	return math.Sqrt(math.Pow(v.X, 2) + math.Pow(v.Y, 2))
}

// Scale method of type Vertex
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	myCar := car{brand: "Audi", price: 40000}
	changeCar(myCar, "Renault", 20000)
	fmt.Println(myCar)

	myCar.changeCar("Renault", 20000)
	fmt.Println(myCar)

	// & will be added implicitly
	myCar.changeCarWithPointer("Renault", 20000)
	fmt.Println(myCar)

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)

	// Valid ways
	yourCar.changeCarWithPointer("Vw", 30000)
	fmt.Println(myCar)

	// ex2
	v := Vertex{3., 4.}
	fmt.Println(v.Abs())

	v.Scale(2)
	fmt.Println(v.Abs())
}
