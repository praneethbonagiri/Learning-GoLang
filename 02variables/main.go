package main

import "fmt"

// declaring a variable ourside a method we cannot use walrus operator

const privateVar string = "This is a private variable so we have written in camel case"
const PublicVar string = "This is a public variable so we have written in pascal case"

// no need to declare any access modifiers like in java
// this pascal case is applicable for variables and functions also

func main() {
	var username string = "Praneeth"
	fmt.Println("Hello " + username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedin bool = true
	fmt.Println(isLoggedin)
	fmt.Printf("Variable is of type: %T \n", isLoggedin)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type: %T \n", smallVal)

	var smallFloat float32 = 255.00127432345046
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type: %T \n", smallFloat)

	var bigFloat float64 = 255.00127432345046
	fmt.Println(bigFloat)
	fmt.Printf("Variable is of type: %T \n", bigFloat)

	// default values and aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of type: %T \n", anotherVariable)

	var anotherString string
	fmt.Println(anotherString)
	fmt.Printf("Variable is of type: %T \n", anotherString)

	// implicit declaration
	// := walrus operator
	number := 3875
	fmt.Println(number)
	fmt.Printf("Variable is of type: %T \n", number)

	fmt.Println(PublicVar)
	fmt.Print(privateVar)
}
