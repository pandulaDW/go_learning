package main

import "fmt"

// Celsius(t) and Fahrenheit(t) are conversions, not function calls. They don’t
// change the value or representation in any way, but they make the change of meaning explicit.

// For every type T, there is a corresponding conversion operation T(x) that converts the value x
// to type T.

// A conversion from one type to another is allowed if both have the same underlying
// type, or if both are unnamed pointer types that point to variables of the same underlying type;

// Celsius type
type Celsius float64

// String method defined on type Celsius
func (c Celsius) String() string {
	return fmt.Sprintf("%g'C", c)
}

// Fahrenheit type
type Fahrenheit float64

// Temperature constants
const (
	AbsoluteZero Celsius = -273.15
	FreezingC    Celsius = 0
	BoilingC     Celsius = 100
)

// CtoF converts a Celsius temperature to Fahrenheit.
func CtoF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FtoC converts F to C
func FtoC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

// Comparison operators like == and < can also be used to compare a value of a named typ e to
// another of the same typ e, or to a value of the underlying typ e. But two values of different
// named types cannot be compared directly

// TypeComparison just compare types
func TypeComparison() {
	var c Celsius
	var f Fahrenheit
	fmt.Println(c == 0) // true
	fmt.Println(f >= 0) // true
	// fmt.Println(c == f) // compile error: type mismatch
	fmt.Println(c == Celsius(f)) // true
}

// Many types declare a String method of this form because it controls how values of the typ e
// appear when printed as a string by the fmt package
