package main

func f() {}

func echo(s string) {
	println("echo: " + s)
}

func sum(i, j int) int {
	return i + j
}

func swap(i, j int) (int, int) {
	return j, i
}

func main() {
	f()
	println(f)

	echo("func")

	println(sum(1, 2))
	println(swap(1, 0))
}
