package main

import "fmt"

// Global scope
var (
	a = 12
	b = 14
)

// local scope
func add(x int, y int) {
	z := x + y
	fmt.Println("Sum is :", z)

}

// local scope
func main() {
	x := 10
	y := 20

	add(x, y)
	add(a, b)
	add(x, z) // This will cause an error because z is not defined in this scope

}

// Scope?
// -> The boundaries of elements in a program.

// If something declared out side of main function, its filled in global area. It can be access from anywhere.

// There are 3 times of scopes:
// Global Scope: Variables declared outside any function or block, accessible throughout the package and other packages if exported (capitalized).

// Local Scope: includes variables declared inside functions and inside blocks (curly braces). Variables in blocks are only accessible within those blocks.

// Package Scope: Variables declared outside functions but inside a package (not exported), accessible anywhere within the same package.
