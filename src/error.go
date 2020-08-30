package main

import (
	"log"
	"os"
)

func main() {
	_, err := os.Open("")
	if err != nil {
		log.Fatal(err)
	}
}
