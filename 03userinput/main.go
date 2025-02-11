package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Welcome to user input program")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your age: ")

	// comma ok or comma error syntax
	// there is no try catch like syntax here
	// when we do some operation either it will give some result
	// or it will give us error

	// if we don't care about the error we store it in _
	// everything is a type in Go (like object in Java) even errors
	age, _ := reader.ReadString('\n')
	fmt.Println("Your age is :", age)

	fmt.Printf("The type of variable is %T", age)
}
