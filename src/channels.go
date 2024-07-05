package main

func main() {
	ch := make(chan string)
	go func(chan string) {
		ch <- "Hello, Go!" // write to channel
	}(ch)
	println(<-ch) // read from channel
}
