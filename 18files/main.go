package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	// ideally when we open a file to read data at the end we need to close the file
	// for it we will use a defer statement :)
	// we will use io utils package to work with files

	// let us try to write the below data into a file
	content := "This data is written into the file by using go lang."

	// we will use os module to create files
	// create file in curr directory and passing the file name
	myFile, err := os.Create("./fileByGolang.txt")
	if err != nil {
		panic(err)
	}

	// write string method returns the length of data written into file
	FileLength, err := io.WriteString(myFile, content)
	if err != nil {
		panic(err)
	}

	fmt.Println("The length of data written to file is:", FileLength)

	defer myFile.Close()

	ReadFromFile(myFile.Name())
	// just trying to see how errors print in golang
	// not very long stack trace
	//ReadFromFile("./fileThatDoesNotExists.txt")
}

func ReadFromFile(myFile string) {

	// we will use os package to read from files and directories
	dataBytes, err := os.ReadFile(myFile)
	if err != nil {
		panic(err)
	}

	// byte array not in string
	fmt.Println(dataBytes)
	fmt.Println("The data read from file is\n", string(dataBytes))
}
