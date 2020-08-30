package main

const (
	STR = "string"
	_   = iota // ignore first value
	TWO
	SIX = iota * TWO
	EIGHT
	TEN
)

func main() {
	println(STR, TWO, SIX, EIGHT, TEN)
}
