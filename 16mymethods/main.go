package main

import "fmt"

func main() {
	fmt.Println("Welcome to methods in golang")

	// methods are same as functions but instead of being a normal function
	// they will be a part of struct

	praneeth := User{"Praneeth", "praneeth@go.dev", 25, true}
	fmt.Println("The name of the user is :", praneeth.Name)

	// to get details we need to use the variable name and property each time
	// else we can use methods
	praneeth.PrintStatus()

	fmt.Printf("The name of user is %v and email is %v\n", praneeth.Name, praneeth.Email)
	// changing the mail
	praneeth.SetDummyMail()
	fmt.Printf("The name of user is %v and email is %v\n", praneeth.Name, praneeth.Email)
	// see so if you pass a struc to a method only a copy will be passed
	// instead of the original object
	// so we need to use pointers here

}

type User struct {
	Name   string
	Email  string
	Age    int
	Status bool
}

func (u User) PrintStatus() {
	fmt.Printf("Is the user %v logged in : %v \n", u.Name, u.Status)
}

// just like a setter
func (u User) SetDummyMail() {
	u.Email = "test@go.dev"
	fmt.Println("The dummy mail id is : ", u.Email)
}
