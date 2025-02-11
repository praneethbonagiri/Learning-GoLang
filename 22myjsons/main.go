package main

import (
	"encoding/json"
	"fmt"
)

// this is the general form of struct
// type course struct {
// 	Name     string
// 	Price    int
// 	Platform string
// 	Password string
// 	Tags     []string
// }

// but if we want to use this struct in api go lang provides json package
// to handle it line changing the names just as aliases
// and option to add this to response or ignore
type course struct {
	Name     string   `json:"coursename"`
	Price    int      // we can use the same name if needed
	Platform string   `json:"website"`
	Password string   `json:"-"` // if we want this field to not be a part of response
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Welcome to JSON's in golang")

	// learning encoding of json
	// we just have some data in any type eg: course type
	// we need to store it or convert it into json
	// that is known as encoding of json
	EncodeJson()
	// The problem with this approach was we were
	// able to see all fields like password also
	// and the names of fields are exposed to others which is not a good practice

	// To deocd josn
	// to write data from a response to any object / type
	DecodeJson()
}

func EncodeJson() {

	myCources := []course{
		{"React JS", 299, "learncodeonline.in", "jabfd8724", []string{"web-dev", "js", "frontend"}},
		{"MERN", 199, "learncodeonline.in", "189fhnvw", []string{"DB", "fullstack"}},
		{"Angular", 299, "learncodeonline.in", "nqdkj289r", nil}, // no tags
	}

	// finalJson, err := json.Marshal(myCources)
	// here the data is printed as a single line
	// to better indent we can use another method
	finalJson, err := json.MarshalIndent(myCources, "", "\t")
	// better intendation, better redability :)

	if err != nil {
		panic(err)
	}

	// fmt.Println(finalJson) // prints byte data as above method returns bytes
	fmt.Printf("%s\n", finalJson)
}

func DecodeJson() {

	jsonDataFromWebBytes := []byte(`
	{
                "coursename": "MERN",
                "Price": 199,
                "website": "learncodeonline.in",
                "tags": ["DB","fullstack"]
        }
	`)

	var newCourse course
	err := json.Unmarshal(jsonDataFromWebBytes, &newCourse)

	if err != nil {
		fmt.Println("Json is not valid")
	} else {
		fmt.Println("Json is valid")
		// fmt.Println(newCourse)
		// to better format we use as below
		// it prints both name and values instead of just values
		fmt.Printf("%#v\n", newCourse)
	}

	// if we want the data to be stored as key value pairs instead of a struct
	// we can use as below
	var newCourseValues map[string]interface{}
	// key is string
	// value is interface because we can get values as int, float, string, slice and another struct
	json.Unmarshal(jsonDataFromWebBytes, &newCourseValues)
	fmt.Printf("%#v\n", newCourseValues)

	// looping through for loop
	for key, val := range newCourseValues {
		fmt.Printf("%v --> %v\n", key, val)
	}

}
