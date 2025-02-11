package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to switch case in golang")

	var favNum int = 11

	switch favNum {
	case 1:
		fmt.Println("Ohh your favourite number is the starting number 1")
	case 2:
		fmt.Println("Your favourite number is smallest even number")
	case 11:
		fmt.Println("Your favourite number is 11")
	default:
		fmt.Println("Wow that's a unique favourite number we have seen so far")
	}

	// in other languages we will write break statements right but in go they are not required
	// in go we can consider break by default
	// if we need to execute next cases we can use fallthrough

	// lets us try to gerenate random numbers and use switch case
	// mimic ludo dice, numbers from 1 to 6. But rand gives 0 to 5 so add 1
	diceNumber := rand.Intn(6) + 1

	switch diceNumber {
	case 1:
		fmt.Println("Ha ha move only 1 step")
	case 2:
		fmt.Println("Move only 2 steps")
	case 3:
		fmt.Println("Move only 3 steps")
	case 4:
		fmt.Println("Great move 4 steps")
		fallthrough
	case 5:
		fmt.Println("Wow move 5 steps")
	case 6:
		fmt.Println("You have done it, move 6 steps and roll again")

	}

}
