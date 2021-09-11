# Hello Go
My place to collect information about Go

- [Hello Go](#hello-go)
  - [How to view slides](#how-to-view-slides)
  - [Main function](#main-function)
  - [Keywords](#keywords)

## How to view slides
https://go-talks.appspot.com/github.com/marcin-janas/hello-go/hello-go.slide#1

## Main function
```go
package main

func main() {
	println("Hello, Go!") // Built-in println function
}
```
- [ref/spec#Program_execution](https://golang.org/ref/spec#Program_execution)
- [ref/spec#Bootstrapping](https://golang.org/ref/spec#Bootstrapping)
- [doc/effective_go.html#package-names](https://golang.org/doc/effective_go.html#package-names)
- [tour/basics/1](https://tour.golang.org/basics/1)
- [package/builtin/#println](https://golang.org/pkg/builtin/#println)

## Keywords
Go only has *25* keywords!
```go
    break, case, chan, const, continue,
    default, defer, else, fallthrough, for,
    func, go, goto, if, import,
    interface, map, package, range, return,
    select, struct, switch, type, var
```

"Syntax is clean and light on keywords"

- [ref/spec#Keywords](https://golang.org/ref/spec#Keywords)
- [src/go/.../token.go#L97](https://github.com/golang/go/blob/master/src/go/token/token.go#L97)
- [doc/faq#principles](https://golang.org/doc/faq#principles)
