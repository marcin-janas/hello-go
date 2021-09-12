# Hello Go
My place to collect information about Go

- [Hello Go](#hello-go)
  - [How to view slides](#how-to-view-slides)
  - [Main function](#main-function)
  - [Keywords](#keywords)
  - [Built-in functions](#built-in-functions)
  - [Basic types](#basic-types)
  - [Zero values](#zero-values)
  - [Constants](#constants)
  - [Variables](#variables)
  - [Import](#import)
  - [Init](#init)
  - [Blank identifier](#blank-identifier)
  - [Unused](#unused)
  - [If](#if)
  - [For](#for)
  - [Switch](#switch)

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
- [doc/effective_go#package-names](https://golang.org/doc/effective_go.html#package-names)
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

* Constants
```go
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
```
- [ref/spec#Constants](https://golang.org/ref/spec#Constants)
- [ref/spec#Iota](https://golang.org/ref/spec#Iota)
- [doc/effective_go#constants](https://golang.org/doc/effective_go.html#constants)

## Variables
```go
package main

var (
	str1 = "test"
	int1 int
)

func main() {
	var str2 string
	int2 := 2
	println(str1, int1, str2, int2)
}
```
- [ref/spec#Variables](https://golang.org/ref/spec#Variables)
- [doc/effective_go#variables](https://golang.org/doc/effective_go.html#variables)

## Import
```go
package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("Hello, Go!") // Println function from the fmt package
	log.Println("Hello, Go!") // Println function from the log package
}
```
- [doc/effective_go#package-names](https://golang.org/doc/effective_go.html#package-names)
- [tour/basics/2](https://tour.golang.org/basics/2)
- [pkg/fmt/#Println](https://pkg.go.dev/fmt#Println)
- [pkg/log/#Println](https://pkg.go.dev/log#Println)

## Init
```go
package main

func init() {
	println("init")
}

func main() {
	println("main")
}
```
- [doc/effective_go#init](https://golang.org/doc/effective_go.html#init)
- [ref/spec#Package_initialization](https://golang.org/ref/spec#Package_initialization)

## Blank identifier
```go
package main

import _ "fmt"

func main() {
	str := "test"
	for _, r := range str {
		_ = r
	}
}
```
- [ref/spec#Blank_identifier](https://golang.org/ref/spec#Blank_identifier)
- [doc/effective_go#blank](https://golang.org/doc/effective_go.html#blank)

## Unused
```go
package main

import "fmt"

func main() {
	msg := "Hello, Go!"
}
```
- [doc/effective_go#blank_unused](https://golang.org/doc/effective_go.html#blank_unused)
- [doc/faq#unused_variables_and_imports](https://golang.org/doc/faq#unused_variables_and_imports)

## If
"A language needs only one conditional control flow construct"
```go
package main

func main() {
	a := 0
	if a > 0 {
		println(">")
	} else if a < 0 {
		println("<")
	} else {
		println("=")
	}
}
```
- [ref/spec#If_statements](https://golang.org/ref/spec#If_statements)
- [doc/effective_go.html#if](https://golang.org/doc/effective_go.html#if)
- [doc/faq#Does_Go_have_a_ternary_form](https://golang.org/doc/faq#Does_Go_have_a_ternary_form)

## For
"Go has only one looping construct, the for loop"
```go
package main

func main() {
	for {
	}

	s := []int{1, 2, 3, 4, 5}
	for n, v := range s {
		println(n, v)
	}

	for i := 0; i <= 9; i++ {
		println(i)
	}
}
```
- [ref/spec#For_statements](https://golang.org/ref/spec#For_statements)
- [doc/effective_go.html#for](https://golang.org/doc/effective_go.html#for)
- [tour/flowcontrol/1](https://tour.golang.org/flowcontrol/1)

## Switch
"A switch statement is a shorter way to write a sequence of if - else statements"
```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	r := rand.Intn(3)
	switch r {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	default:
		fmt.Println(2)
	}
}
```
- [ref/spec#Switch_statements](https://golang.org/ref/spec#Switch_statements)
- [doc/effective_go#switch](https://golang.org/doc/effective_go#switch)
- [tour/flowcontrol/9](https://tour.golang.org/flowcontrol/9)
- [pkg/math/rand#Seed](https://pkg.go.dev/math/rand#Seed)
- [pkg/math/rand#Intn](https://pkg.go.dev/math/rand#Intn)

