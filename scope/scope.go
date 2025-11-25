package main

import (
	"fmt"

	"example.com/mathlib"
)

// Global scope
var (
	a = 12
	b = 14
)

// local scope
// func add(x int, y int) {
// 	z := x + y
// 	fmt.Println("Sum is :", z)

// }

// local scope
func main() {
	// x := 10
	// y := 20

	// add(x, y)
	// add(a, b)
	// add(x, z) // This will cause an error because z is not defined in this scope

	fmt.Println("Custom package mathlib")
	mathlib.Add(10, 20)

}
