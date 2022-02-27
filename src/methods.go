package main

import "fmt"

type rectangle struct {
	description string
	a, b        int
}

func (r rectangle) area() int {
	return r.a * r.b
}

func main() {
	r := rectangle{"square", 2, 2}
	fmt.Println(r.area())
}
