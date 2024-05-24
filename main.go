package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Welcome to make server project")
	// fileserver := http.FileServer(http.Dir("./static"))
	// http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting the server on port 4000")
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal(err)
	}

}

func formHandler(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
		fmt.Println("Error")
	}

	fmt.Fprint(w, "POST request successful")

	username := r.FormValue("username")
	password := r.FormValue("password")
    
	fmt.Fprint(w, password)
	fmt.Fprint(w, username)
	
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// CHECK1
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}

	//CHECK2
	if r.Method != "GET" {
		http.Error(w, "Method is not supported", http.StatusNotFound)
		return
	}

	fmt.Fprint(w, "Hello from Kushal Jindal")
}
