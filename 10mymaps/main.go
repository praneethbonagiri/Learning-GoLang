package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome to maps in golang")

	coursesMap := make(map[string]string)

	coursesMap["JS"] = "Java Script"
	coursesMap["PY"] = "Python"
	coursesMap["J"] = "Java"
	coursesMap["Ru"] = "Ruby"

	fmt.Println("List of all courses - ", coursesMap)
	fmt.Println("Full course for key JS is ", coursesMap["JS"])

	delete(coursesMap, "Ru")
	fmt.Println("List of all courses - ", coursesMap)

	// loop through map

	for key, value := range coursesMap {
		fmt.Printf("For the key %v the value is %v\n", key, value)
	}
}
