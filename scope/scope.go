package main

import "fmt"

// Global scope
var (
	a = 12
	b = 14
)

func add(x int, y int) {
	z := x + y
	fmt.Println("Sum is :", z)

}

func main() {
	x := 10
	y := 20

	add(x, y)
	add(a, b)
	add(x, z) // This will cause an error because z is not defined in this scope

}
