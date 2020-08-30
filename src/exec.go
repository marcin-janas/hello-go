package main

import (
	"log"
	"os/exec"
)

func main() {
	output, err := exec.Command("echo", "Hello Go!").Output()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("output: %s", output)
}
