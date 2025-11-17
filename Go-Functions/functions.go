package main

import "fmt"

func add(num1 int, num2 int) int {
	sum := num1 + num2
	return sum

}

func main() {
	a := 2
	b := 3

	result := add(a, b)
	fmt.Println("The sum is:", result)
}