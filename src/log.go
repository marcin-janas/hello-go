package main

import "log"

func f1() {
	log.SetPrefix("f1:")
	log.Println("some message")
}

func f2() {
	log.SetPrefix("f2:")
	log.Println("some message")
}

func main() {
	log.SetFlags(0)
	f1()
	f2()
}
