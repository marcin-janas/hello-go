package main

import "fmt"

type Test struct {
	test string
}

func main() {
	test := Test{"test"}
	fmt.Println(test.test)
}
