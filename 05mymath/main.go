package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Lets explore Math and rand package in golang")

	// var num1 int = 100
	// var num2 float64 = 3.14687
	// fmt.Println(num1 + num2)
	// this is incorrect as we connot directly add int and float

	// random number in math/rand package
	// myRandNum := rand.Intn(100)
	// fmt.Println(myRandNum)

	// random number in crypto/rand package
	myRandNum2, _ := rand.Int(rand.Reader, big.NewInt(100))
	fmt.Println(myRandNum2)

	for i := 0; i < 5; i++ {
		fmt.Println(rand.Int(rand.Reader, big.NewInt(10)))
	}

}
