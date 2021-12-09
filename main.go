package main

import (
	"fmt"
	"net/http"
)

// test de git commit dgil.
func main() {

	http.HandleFunc("/API/", helloHandler)

	http.ListenAndServe(":8080", nil)

}

func helloHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Hello World GOOOOOOOOOOOO/n")
}
