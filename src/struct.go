package main

import "fmt"

type Test struct {
	test *string
}

func main() {
	test := Test{}
	fmt.Println(*test.test)
}
