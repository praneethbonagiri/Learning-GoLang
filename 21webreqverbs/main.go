package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to web requests in go lang, like GET POST")
	// uncomment below line to run get request
	// PerformGetReq()

	// uncomment below line to run post req where data is as json
	// PerformPostJsonReq()

	// uncomment below line to run post req where data is as form
	PerformPostFormReq()
}

// function for get request
func PerformGetReq() {
	const myUrl = "http://localhost:8000/get"

	resp, err := http.Get(myUrl)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()
	fmt.Println("The status code of response is :", resp.StatusCode)
	fmt.Println("The length of response is :", resp.ContentLength)

	dataBytes, _ := io.ReadAll(resp.Body)
	// it works but lets try another way
	//fmt.Println("The data is : ", string(dataBytes))
	var responseData strings.Builder

	dataLen, _ := responseData.Write(dataBytes)
	fmt.Println("The length of data is :", dataLen)
	fmt.Println(responseData.String())
}

// function for post request using json data
func PerformPostJsonReq() {
	const myUrl = "http://localhost:8000/post"

	requestBody := strings.NewReader(`
		{
			"coursename" : "Golang",
			"price" : 0,
			"platform" : "youtube"
		}
	`)

	response, err := http.Post(myUrl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	responseDataBytes, _ := io.ReadAll(response.Body)
	var responseData strings.Builder
	responseData.Write(responseDataBytes)
	fmt.Println("Response data is :\n", responseData.String())
}

// function for post request with form data
func PerformPostFormReq() {
	const myUrl = "http://localhost:8000/postform"

	// formdata - data as key and value pairs
	formdata := url.Values{}

	formdata.Add("firstname", "Praneeth")
	formdata.Add("lastname", "Bonagiri")
	formdata.Add("email", "praneeth@go.dev")

	response, err := http.PostForm(myUrl, formdata)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	responseBytes, _ := io.ReadAll(response.Body)
	responseData := string(responseBytes)

	fmt.Println("The response from postform url is: \n", responseData)

}
