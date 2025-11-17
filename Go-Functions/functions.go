package main

import "fmt"

func getNumbers(num1 int, num2 int) (int, int) {

	sum := num1 + num2
	mul := num1 * num2
	return sum, mul
}

func main() {
	a := 10
	b := 15

	sum, mul := getNumbers(a, b)
	fmt.Println("Sum:", sum)
	fmt.Println("Multiplication:", mul)

}