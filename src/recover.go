package main

import "sync"

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go f()
	wg.Wait()
}

// START OMIT
func f() {
	defer func() {
		if recover() != nil {
			println("recover f")
		}
		println("defer f")
		wg.Done()
	}()
	p()
}

func p() {
	defer println("defer p")
	panic("panic p")
}

// END OMIT
