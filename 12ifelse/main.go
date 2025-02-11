package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to if else in golang")

	numUsers := 11
	var message string

	if numUsers < 10 {
		message = "Normal Usage"
	} else if numUsers > 10 {
		message = "Look out!.. Server can creash :)"
	} else {
		message = "Exactly 10 users"
	}

	fmt.Println("Number of users are : ", message)

	// fmt.Println("Let's play a game \n Give us a number and we will tell you if its an even or odd number :")

	// reader := bufio.NewReader(os.Stdin)

	// number, _ := reader.ReadString('\n')
	// intNum, _ := strconv.ParseInt(strings.TrimSpace(number), 10, 32)

	// if intNum%2 == 0 {
	// 	fmt.Println("You have entered an even number")
	// } else {
	// 	fmt.Println("Ooh you have entered an odd number")
	// }

	// we can also initialize variables in if statement
	// generally it will be used at time of any api calls kind of situation
	if num := 3; num < 0 {
		fmt.Println("Negative number :(")
	} else {
		fmt.Println("Positive number :)")
	}

}
