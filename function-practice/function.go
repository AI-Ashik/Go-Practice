package main

import "fmt"

func welcomeMessage() {

	fmt.Println("Welcome to the application")
}

func getName() string {
	var name string
	fmt.Println("Enter your name : ")
	fmt.Scanln(&name)
	fmt.Println("Welcome : ", name)
	return name
}

func getNumbers() (int, int) {
	var num1 int
	var num2 int
	fmt.Println("Enter numbe 1 : ")
	fmt.Scanln(&num1)
	fmt.Println("Enter numbe 2 : ")
	fmt.Scanln(&num2)
	return num1, num2
}

func summation(num1 int, num2 int) int {

	sum := num1 + num2
	return sum
}

func totalSum(sum int) {

	fmt.Println("Sum of two numbers is : ", sum)

}

func finalMessage(name string) {
	fmt.Println("Thank you for using the appliocation", name)

}

func main() {
	welcomeMessage()
	name := getName()
	num1, num2 := getNumbers()
	sum := summation(num1, num2)
	totalSum(sum)
	finalMessage(name)

}
