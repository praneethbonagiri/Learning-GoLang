package main

import "fmt"

func main() {
	fmt.Println("Welcome to structs in golang")

	// structs in golang are just like classes in java
	// things to note are golang does not support inheritance
	// no concept of super and parent

	// creating a variable with our struct
	praneeth := User{"Praneeth", "praneeth@go.dev", 25, true}
	fmt.Println(praneeth)

	// if we want output in more detailed manner we can use
	fmt.Printf("The user details are - %+v\n", praneeth)

	// to print any individual details
	fmt.Printf("The name of user is %v and email id is %v", praneeth.Name, praneeth.Email)
}

// creating a struct named User (like a class in java)
// see here we are following pascal case ie first letter is capital
// so the struct and its variables are public
type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

// so type and struct key words we used here
// anything basically in go is a type (just as an object in java)
