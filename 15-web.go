package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}
func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	fmt.Println("hola!")

	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)

	fmt.Println("Server starting ...")
	http.ListenAndServe(":3000", nil)

}

/*
After running this code,
- go to "http://localhost:3000/" in browser, and we will get "Hello World"
- go to "http://localhost:3000/about", and we will get "About"
*/
