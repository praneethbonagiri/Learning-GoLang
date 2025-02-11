package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// model for course - generally this goes into a separate file
type Course struct {
	CourseId   string  `json:"courseid"`
	CourseName string  `json:"coursename"`
	Price      int     `json:"price"`
	Author     *Author `json:"author"`
}

type Author struct {
	Fullname string `json:"fullname"`
	Website  string `json:"website"`
}

// fake db
var courses []Course

// helper functions can be in a separate file
func (c *Course) IsEmpty() bool {
	return c.CourseName == ""
}

func main() {
	fmt.Println("Welcome to API's in golang, here we will use slice instead of actual db to store data")

	router := mux.NewRouter()

	// seeding or adding some data to slice as a fake db
	courses = append(courses, Course{CourseId: "2", CourseName: "React JS", Price: 199, Author: &Author{Fullname: "Praneeth", Website: "youtube.com"}})
	courses = append(courses, Course{CourseId: "4", CourseName: "Golang", Price: 0, Author: &Author{Fullname: "Hitesh Choudhary", Website: "learncodeonline.in"}})

	// adding routing and handlers or controllers
	router.HandleFunc("/", serveHome).Methods("GET")
	router.HandleFunc("/courses", getAllCourses).Methods("GET")
	router.HandleFunc("/course/{id}", getOneCourse).Methods("GET")
	router.HandleFunc("/course", createOneCourse).Methods("POST")
	router.HandleFunc("/course/{id}", updateOneCourse).Methods("PUT")
	router.HandleFunc("/course/{id}", deleteOneCourse).Methods("DELETE")

	// listen on a port
	log.Fatal(http.ListenAndServe(":4000", router))

}

// controllers can be in different files

// serve home route
func serveHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to API's in golang</h1>"))
}

func getAllCourses(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get all courses")
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(courses)
}

func getOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get one course")
	w.Header().Set("Content-Type", "application/json")

	// grab params from request
	params := mux.Vars(r)

	// loop through courses and match id which user has sent and return it
	for _, course := range courses {
		if course.CourseId == params["id"] {
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id")
	return
}

func createOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Create one course")
	w.Header().Set("Content-Type", "application/json")

	// what if body is empty
	if r.Body == nil {
		json.NewEncoder(w).Encode("Request body should not be empty, send some data")
	}

	// what if data is like {}

	var newCourse Course
	_ = json.NewDecoder(r.Body).Decode(&newCourse)

	if newCourse.IsEmpty() {
		json.NewEncoder(w).Encode("Course name cannot be null")
		return
	}

	// generate a unique id, convert into string and then add into slice
	intId := rand.Intn(100)
	newCourse.CourseId = strconv.Itoa(intId)

	courses = append(courses, newCourse)
	json.NewEncoder(w).Encode(newCourse)
	return
}

func updateOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Update one course")
	w.Header().Set("Content-Type", "application/json")

	// first get id from request params
	params := mux.Vars(r)

	// find this index in courses slices then remove it and then add new course
	for index, course := range courses {
		if course.CourseId == params["id"] {
			courses = append(courses[:index], courses[index+1:]...)

			var modifiedCourse Course
			json.NewDecoder(r.Body).Decode(&modifiedCourse)
			modifiedCourse.CourseId = params["id"]

			courses = append(courses, modifiedCourse)
			json.NewEncoder(w).Encode(modifiedCourse)
			return
		}
	}

	json.NewEncoder(w).Encode("No course found with given id unable to update")
}

func deleteOneCourse(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Delete one course")
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	deleteCourseWithId := params["id"]

	for index, course := range courses {
		if course.CourseId == deleteCourseWithId {
			courses = append(courses[:index], courses[index+1:]...)
			json.NewEncoder(w).Encode(course)
			return
		}
	}

	json.NewEncoder(w).Encode("No course was found with given id, unable to delete")
	return
}
