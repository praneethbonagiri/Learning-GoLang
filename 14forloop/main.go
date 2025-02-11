package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang, there is only for loop in go nothing else")

	// loops, break, continue

	words := []string{"Hello", "How", "When", "Like", "Good", "Bye", "Next"}

	// normal for loop
	// looping through index
	for i := 0; i < len(words); i++ {
		fmt.Println(words[i])
	}

	// loop through index
	for idx := range words {
		fmt.Println(words[idx])
	}

	// for each loop
	for _, word := range words {
		fmt.Println(word)
	}

	// first will be index, next will be value
	for idx, word := range words {
		fmt.Printf("The index is %v and the value is %v\n", idx, word)
	}

	// using break, continue and goto
	for idx := range words {
		fmt.Println("Starting for loop idx is ", idx)

		if idx == 2 {
			fmt.Println("Will continue now")
			continue
		}

		if idx == 5 {
			fmt.Println("Will break now")
			break
		}

		if idx == 4 {
			fmt.Println("Will call goto block")
			goto endLine
		}

		fmt.Println("Ending for loop idx is ", idx)
	}

endLine:
	fmt.Println("This is example usage of goto statement in for loop")

	// we can also use for loop as while loop
	currNum := 0

	for currNum < 5 {
		fmt.Println("Current number is :", currNum)
		currNum++
	}
}
