package main

import "fmt"

func main() {
	fmt.Println("Welcome to defer statements in golang")

	// defer statement is used for statements or lines of code which we want to execute at the end of function
	// see the order of execution of below statements
	fmt.Println("First line")
	defer fmt.Println("Second line")
	fmt.Println("Third line")

	// more examples
	fmt.Println("example1")
	defer fmt.Println("example2")
	defer fmt.Println("example3")
	defer fmt.Println("example4")
	fmt.Println("example5")
	// op will be 1, 5, 4, 3, 2

	myDefer()
	fmt.Println()

}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
