package main

import (
	"fmt"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/hello" {
		http.Error(w, "404 error not found", http.StatusNotFound)

		return
	}
	if r.Method != "GET" {
		http.Error(w, "405 Error method not supported", http.StatusMethodNotAllowed)

		return
	}

	fmt.Fprintf(w, "Hello World !")

}

func contactFormHandler(w http.ResponseWriter, r *http.Request) {

	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm error : %v", err)
		return
	}
	fmt.Fprintf(w, "Post request success ! \n")

	name := r.FormValue("name")
	email := r.FormValue("email")

	fmt.Fprintf(w, "Name : %s\n", name)
	fmt.Fprintf(w, "Email : %s\n", email)
}

func main() {

	fileserver := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileserver)

	http.HandleFunc("/contact", contactFormHandler)

	http.HandleFunc("/hello", helloHandler)

	fmt.Printf("Server stared at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
