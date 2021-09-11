# Hello Go
My place to collect information about Go

- [Hello Go](#hello-go)
  - [How to view slides](#how-to-view-slides)
  - [Main function](#main-function)
  - [Keywords](#keywords)
  - [Built-in functions](#built-in-functions)
  - [Basic types](#basic-types)
  - [Zero values](#zero-values)

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
- [pkg/builtin/#println](https://pkg.go.dev/builtin#println)

## Keywords
Go only has **25** keywords!
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

## Built-in functions
There is no guarantee that the **print** and **println** functions will stay in the language!
```go
    append, cap, close, complex, copy,
    delete, imag, len, make, new,
    panic, print, println, real, recover
```
- [ref/spec#Built-in_functions](https://golang.org/ref/spec#Built-in_functions)
- [pkg/builtin](https://pkg.go.dev/builtin)

## Basic types
```go
string
bool
uint, uint8, uint16, uint32, uint64, uintptr, int, int8, int16, int32, int64
byte // alias for uint8
rune // alias for int32
float32, float64
complex64, complex128
pointer, function, interface, slice, channel, map
array, struct
```
- [ref/spec#Types](https://golang.org/ref/spec#Types)
- [tour/basics/11](https://tour.golang.org/basics/11)

## Zero values
```go
"" // for strings
false // for booleans
0 // for numeric types
nil // for pointers, functions, interfaces, slices, channels, and maps
// each element of an array of structs will have its fields zeroed if no value is specified
```
- [ref/spec#The_zero_value](https://golang.org/ref/spec#The_zero_value)
- [tour/basics/12](https://tour.golang.org/basics/12)
