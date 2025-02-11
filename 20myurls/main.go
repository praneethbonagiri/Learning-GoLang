package main

import (
	"fmt"
	"net/url"
)

const myUrl = "https://lco.dev:3000/learn?coursename=golang&paymentid=mdzncowu4tu423gbl"

func main() {
	fmt.Println("Welcome to urls in go lang")

	fmt.Println("The url is ", myUrl)

	// parsing a url
	parsedUrl, _ := url.Parse(myUrl)

	fmt.Printf("The type of parsed url is %T\n", parsedUrl)
	fmt.Println(parsedUrl.Scheme)
	fmt.Println(parsedUrl.Host)
	fmt.Println(parsedUrl.Path)
	fmt.Println(parsedUrl.Hostname())
	fmt.Println(parsedUrl.Port())
	fmt.Println(parsedUrl.RawQuery) // params as string all together

	params := parsedUrl.Query()
	fmt.Printf("the type of params is %T\n", params)
	fmt.Println(params) // params in from of map

	for key, val := range params {
		fmt.Printf("The param is %v and value is: %v\n", key, val)
	}

	// constructing a url from parts
	urlParts := &url.URL{
		Scheme:   "https",
		Host:     "lco.dev",
		Path:     "/tutcss",
		RawQuery: "user=praneeth",
	}

	anotherUrl := urlParts.String()
	fmt.Println(anotherUrl)
}
