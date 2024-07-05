package main

func add[T int | string](s []T) T {
	var r T

	for _, v := range s {
		r += v
	}

	return r
}

func main() {
	i := []int{1, 2, 3}
	s := []string{"One", "Two", "Three"}
	println(add(i))
	println(add(s))
}
