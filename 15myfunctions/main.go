package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to functions in golang")
	greeter()

	result := normalAdd(3, 5)
	fmt.Println(result)

	// we cannot write a new function inside another function
	// function with variable number of arguments
	proResult := proAdd(3, 5, 9, 6, 22)
	fmt.Println("Pro result is :", proResult)

	// function with multiple return values
	// make sure to use the walarus operator
	num, message := multipleReturnValues(6)

	fmt.Printf("The number is %v, and the messages is: %v", num, message)
}

func greeter() {
	fmt.Println("Hello from function to golang")
}

func normalAdd(num1 int, num2 int) int {
	return num1 + num2
}

func proAdd(values ...int) int {
	sum := 0

	for _, value := range values {
		sum += value
	}

	return sum
}

func multipleReturnValues(num int) (int, string) {
	return num + 5, "Thanks for coding"
}
