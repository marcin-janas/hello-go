package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// START OMIT
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, Go!")
	})
	log.Fatal(http.ListenAndServe(":8080", nil))
	// END OMIT
}
