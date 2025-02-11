package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in go")

	// slices are arrays only under hood
	// we can consider slices as if add one layer of abstraction over arrays

	var fruitList = []string{"Apple", "Guava", "Berries", "Straw berries"}
	fmt.Printf("Type of fruit list is %T\n", fruitList)

	// to add more elemets we need to use append method
	fruitList = append(fruitList, "Mango", "Grapes", "Blue berries")
	fmt.Println(fruitList)

	// we can slice the array based on index
	// just like in python
	fruitList = append(fruitList[:6])
	fmt.Println(fruitList)

	fruitList = append(fruitList[2:])
	fmt.Println(fruitList)

	highScores := make([]int, 5)

	highScores[0] = 750
	highScores[1] = 540
	highScores[2] = 399
	highScores[3] = 994
	highScores[4] = 327

	fmt.Println(highScores)
	fmt.Println("Length if high scores is ", len(highScores))

	// if we un comment these we will get error
	// panic: runtime error: index out of range [5] with length 5
	// since we have declared the length to be 5
	// we cannot increase the length directly
	// highScores[5] = 873
	// fmt.Println(highScores)

	// to increase the length ie to add extra elements we need to use append method only

	highScores = append(highScores, 239, 448, 777, 634)
	fmt.Println(highScores)
	fmt.Println("Length if high scores is ", len(highScores))

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("Sorted slice is")
	fmt.Println(highScores)

	fmt.Println(sort.IntsAreSorted(highScores))

	// removing a value from slices based on index
	// we can only remove using append method

	var courseList = []string{"Java", "Golang", "Java script", "Python", "Dot net", "C++", "Ruby", "Rust"}
	fmt.Println(courseList)

	// delete dot net from the slice
	// so the index of dot net is 4
	var index int = 4
	courseList = append(courseList[:index], courseList[index+1:]...)
	fmt.Println(courseList)
	// so we use slice... to get all the elements insead of a slice
	// so the append method takes 2 arguments
	// first a slice
	// next var arg array like in java any number of arguments not fixed
	// that's why we can pass courseList[index+1:]...

}
