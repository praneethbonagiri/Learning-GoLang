package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study in golang")

	presentTime := time.Now()
	fmt.Println(presentTime)

	// we need to use the same mentioned date and time place holders
	// like 01-02-2006 for date
	fmt.Println(presentTime.Format("01-02-2006"))                 // mm-dd-yyyy
	fmt.Println(presentTime.Format("01-02-2006 Monday"))          // mm-dd-yyyy week day
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // mm-dd-yyyy HH:MM:SS week day

	createdDate := time.Date(1999, time.October, 11, 2, 53, 0, 0, time.UTC)
	fmt.Println(createdDate)
}
