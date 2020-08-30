package main

import "flag"

func main() {
	f := flag.Bool("test", false, "test description")
	flag.Parse()

	println(*f)
}
