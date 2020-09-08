package main

import (
	"log"
	"net/http"
)

func main() {
	// START OMIT
	resp, err := http.Get("https://go.dev")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(resp.StatusCode)
	// END OMIT
}
