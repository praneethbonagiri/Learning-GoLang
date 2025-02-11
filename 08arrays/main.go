package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in go")

	var fruitList [4]string

	fruitList[0] = "Apple"
	fruitList[1] = "Banana"
	fruitList[3] = "Grapes"

	fmt.Println("fruit list is ", fruitList)
	fmt.Println("length of fruit list is ", len(fruitList))

	var vegList = [10]string{"Potato", "Carrot", "Bottle guard"}
	fmt.Println("Vegetable list is", vegList)
	fmt.Println("length of veg list is", len(vegList))

	// this is considered as a slice instead of array since we have not specified
	// the size at start
	// var vegList2 = []string{"Potato", "Carrot", "Bottle guard"}
	// fmt.Println("Vegetable list 2 is", vegList2)
	// fmt.Println("length of veg list 2 is", len(vegList2))
}
