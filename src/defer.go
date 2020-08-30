package main

func main() {
	println("func main")
	f1()
}

func f1() {
	defer f2()
	println("func f1")
}

func f2() {
	println("func f2")
}
