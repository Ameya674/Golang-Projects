package main 

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.Error(w, "path not found", http.StatusNotFound)
		return
	}

	if r.Method	!= "GET" {
		http.Error(w, "cannot write on this path", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hello and welcome to our server.")
}

func aboutme(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/aboutme" {
		http.Error(w, "wrong page dumbass", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "cant post or write to this page smarty", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Hey y'all, I am Tuco Salamanca and I run a drug cartel. It is fun. A guy named Heisenberg cooks for me. He's good.")
}

func education(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/education" {
		http.Error(w, "why so dumb??", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "only for reading...", http.StatusNotFound)
		return
	}

	fmt.Fprintf(w, "Completed my schooling from Salt Middle School and currently studying at Seasoning City Technological University.")
}

func form(w http.ResponseWriter, r *http.Request) {
	if error := r.ParseForm(); error != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", error)
		return 
	}

	fmt.Fprintf(w, "request successful")
	
	name := r.FormValue("name")
	city := r.FormValue("city")
	state := r.FormValue("state")
	roll := r.FormValue("roll")

	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "City: %s\n", city)
	fmt.Fprintf(w, "State: %s\n", state)
	fmt.Fprintf(w, "Roll: %s\n", roll)
}

func main() {
	server := http.FileServer(http.Dir("./static"))
	http.Handle("/", server)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/aboutme", aboutme)
	http.HandleFunc("/education", education)
	http.HandleFunc("/form", form)

	fmt.Printf("Server running at port: 3000")

	if error := http.ListenAndServe(":3000", nil); error != nil {
		log.Fatal(error)
	}
}