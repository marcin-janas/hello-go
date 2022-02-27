# Hello Go
My place to collect information about Go

- [Hello Go](#hello-go)
  - [Main function](#main-function)
  - [Keywords](#keywords)
  - [Built-in functions](#built-in-functions)
  - [Basic types](#basic-types)
  - [Zero values](#zero-values)
  - [Constants](#constants)
  - [Variables](#variables)
  - [Packages](#packages)
  - [Import](#import)
  - [Init](#init)
  - [Blank identifier](#blank-identifier)
  - [Unused](#unused)
  - [If](#if)
  - [For](#for)
  - [Switch](#switch)
  - [Function](#function)
  - [Struct](#struct)
  - [Defer](#defer)
  - [Error handling](#error-handling)
  - [Panic and recover](#panic-and-recover)

## Main function
```go
package main

func main() {
	println("Hello, Go!") // Built-in println function
}
```
- [ref/spec#Program_execution](https://go.dev/ref/spec#Program_execution)
- [ref/spec#Bootstrapping](https://go.dev/ref/spec#Bootstrapping)
- [doc/effective_go#package-names](https://go.dev/doc/effective_go#package-names)
- [tour/basics/1](https://go.dev/tour/basics/1)
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

- [ref/spec#Keywords](https://go.dev/ref/spec#Keywords)
- [src/go/.../token.go#L97](https://github.com/golang/go/blob/master/src/go/token/token.go#L97)
- [doc/faq#principles](https://go.dev/doc/faq#principles)

## Built-in functions
There is no guarantee that the **print** and **println** functions will stay in the language!
```go
    append, cap, close, complex, copy,
    delete, imag, len, make, new,
    panic, print, println, real, recover
```
- [ref/spec#Built-in_functions](https://go.dev/ref/spec#Built-in_functions)
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
- [ref/spec#Types](https://go.dev/ref/spec#Types)
- [tour/basics/11](https://go.dev/tour/basics/11)

## Zero values
```go
"" // for strings
false // for booleans
0 // for numeric types
nil // for pointers, functions, interfaces, slices, channels, and maps
// each element of an array of structs will have its fields zeroed if no value is specified
```
- [ref/spec#The_zero_value](https://go.dev/ref/spec#The_zero_value)
- [tour/basics/12](https://go.dev/tour/basics/12)

## Constants
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
- [ref/spec#Constants](https://go.dev/ref/spec#Constants)
- [ref/spec#Iota](https://go.dev/ref/spec#Iota)
- [doc/effective_go#constants](https://go.dev/doc/effective_go#constants)

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
- [ref/spec#Variables](https://go.dev/ref/spec#Variables)
- [doc/effective_go#variables](https://go.dev/doc/effective_go#variables)

## Packages
- [ref/spec#Packages](https://go.dev/ref/spec#Packages)
- [tour/basics/1](https://go.dev/tour/basics/1)
- [doc/effective_go#package-names](https://go.dev/doc/effective_go#package-names)
- [avoid-package-names-like-base-util-or-common](https://dave.cheney.net/2019/01/08/avoid-package-names-like-base-util-or-common)
- [pkg#1](https://pkg.go.dev)
- [pkg#2](https://github.com/pkg)

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
- [doc/effective_go#package-names](https://go.dev/doc/effective_go#package-names)
- [tour/basics/2](https://go.dev/tour/basics/2)
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
- [doc/effective_go#init](https://go.dev/doc/effective_go#init)
- [ref/spec#Package_initialization](https://go.dev/ref/spec#Package_initialization)

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
- [ref/spec#Blank_identifier](https://go.dev/ref/spec#Blank_identifier)
- [doc/effective_go#blank](https://go.dev/doc/effective_go#blank)

## Unused
```go
package main

import "fmt"

func main() {
	msg := "Hello, Go!"
}
```
- [doc/effective_go#blank_unused](https://go.dev/doc/effective_go#blank_unused)
- [doc/faq#unused_variables_and_imports](https://go.dev/doc/faq#unused_variables_and_imports)

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
- [ref/spec#If_statements](https://go.dev/ref/spec#If_statements)
- [doc/effective_go#if](https://go.dev/doc/effective_go#if)
- [doc/faq#Does_Go_have_a_ternary_form](https://go.dev/doc/faq#Does_Go_have_a_ternary_form)

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
- [ref/spec#For_statements](https://go.dev/ref/spec#For_statements)
- [doc/effective_go#for](https://go.dev/doc/effective_go#for)
- [tour/flowcontrol/1](https://go.dev/tour/flowcontrol/1)

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
- [ref/spec#Switch_statements](https://go.dev/ref/spec#Switch_statements)
- [doc/effective_go#switch](https://go.dev/doc/effective_go#switch)
- [tour/flowcontrol/9](https://go.dev/tour/flowcontrol/9)
- [pkg/math/rand#Seed](https://pkg.go.dev/math/rand#Seed)
- [pkg/math/rand#Intn](https://pkg.go.dev/math/rand#Intn)

## Function
```go
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
```
- [ref/spec#Function_types](https://go.dev/ref/spec#Function_types)
- [doc/effective_go#functions](https://go.dev/doc/effective_go#functions)
- [doc/faq#Functions_methods](https://go.dev/doc/faq#Functions_methods)
- [tour/basics/4](https://go.dev/tour/basics/4)
- [tour/basics/5](https://go.dev/tour/basics/5)
- [tour/basics/6](https://go.dev/tour/basics/6)

## Struct
```go
package main

import "fmt"

type rectangle struct {
	description string
	a, b        int
}

func main() {
	r := rectangle{"square", 2, 2}
	fmt.Println(r)
}
```
- [ref/spec#Struct_types](https://go.dev/ref/spec#Struct_types)
- [tour/moretypes/5](https://go.dev/tour/moretypes/5)
- [the-empty-struct](https://dave.cheney.net/2014/03/25/the-empty-struct)

## Defer
```go
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
```
- [ref/spec#Defer_statements](https://go.dev/ref/spec#Defer_statements)
- [doc/effective_go#defer](https://go.dev/doc/effective_go#defer)
- [blog/defer-panic-and-recover](https://go.dev/blog/defer-panic-and-recover)

## Error handling
```go
package main

import "os"

func main() {
	if _, err := os.Open(""); err != nil {
		os.Exit(1)
	}
}
```
- [ref/spec#Errors](https://go.dev/ref/spec#Errors)
- [doc/effective_go#errors](https://go.dev/doc/effective_go#errors)
- [doc/faq#exceptions](https://go.dev/doc/faq#exceptions)
- [blog/error-handling-and-go](https://go.dev/blog/error-handling-and-go)
- [blog/errors-are-values](https://go.dev/blog/errors-are-values)
- [pkg/os#Open](https://pkg.go.dev/os#Open)
- [pkg/os#Exit](https://pkg.go.dev/os#Exit)

## Panic and recover
```go
package main

func main() {
	defer func() {
		if msg, ok := recover().(string); ok {
			print("recover: " + msg)
		}
	}()

	panic("panic msg")
}
```
- [doc/effective_go#panic](https://go.dev/doc/effective_go#panic)
- [doc/effective_go#recover](https://go.dev/doc/effective_go#recover)
- [blog/defer-panic-and-recover](https://go.dev/blog/defer-panic-and-recover)
- [ref/spec#Handling_panics](https://go.dev/ref/spec#Handling_panics)
- [ref/spec#Run_time_panics](https://go.dev/ref/spec#Run_time_panics)
- [ref/spec#Type_assertions](https://go.dev/ref/spec#Type_assertions)
- [medium/go-how-does-a-program-recover](https://medium.com/a-journey-with-go/go-how-does-a-program-recover-fbbbf27cc31e)

