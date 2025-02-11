package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to modules in golang")
	greeter()

	// mux is like a popular routing service in golang
	router := mux.NewRouter()
	// adding path and function to handle of request comes to that path
	router.HandleFunc("/", ServeHome).Methods("GET")

	// start a server to listen on post
	log.Fatal(http.ListenAndServe(":4000", router))
}

func greeter() {
	fmt.Println("This is a greeting function in golang")
}

// function to server a request on specific path
func ServeHome(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Learning golang, this is the default page for a server</h1>"))
}

// go mod tidy
// go mod verify
// go mod graph
// go list all
// go list -m all
// go mod why github.com/gorilla/mux
// go mod vendor ==> copies all dependencies from web into a vendor folder in current folder
// if we need to use the dependencies in vendor folder rather than from cache we need to use below command
// go run -mod=vendor main.go
