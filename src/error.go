package main

import "os"

func main() {
	if _, err := os.Open(""); err != nil {
		os.Exit(1)
	}
}
