package main

func main() {
	defer func() {
		if msg, ok := recover().(string); ok {
			print("recover: " + msg)
		}
	}()

	panic("panic msg")
}
