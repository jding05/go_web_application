package main

import (
	"fmt"
	"log"
	"net/http"
)

// http.ResponseWriter value assembles the HTTP server's response,
// by writing to it, we send data to the HTTP client
// [1:] -> create a sub-slice of Path from the 1st character to the end, drop the leading "/"
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	// tell the http package to handle all requests to the web root ("/") w/ handler
	http.HandleFunc("/", handler)

	// ListenAndServe always returns an error, only return when an unexpencted error occur
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// test 
// http://localhost:8080/monkeys
// -> Hi there, I love monkeys!