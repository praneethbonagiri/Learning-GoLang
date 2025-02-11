package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in go")

	myNumber := 99
	fmt.Println("The number is: ", myNumber)

	var myptr *int = &myNumber
	fmt.Println("The pointer is pointing to memory", myptr)
	fmt.Println("The value that is present in the pointer pointing location is, ", *myptr)

	// of we do any operation on pointer the actual value will also be changed
	// because pointer is only pointing to the memory location and not storing actual value

	*myptr = *myptr + 1

	fmt.Println("The value in myNumber after doing operation on pointer is ", myNumber)
}
