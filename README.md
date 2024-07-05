# Hello Go

My place to collect information about Go

- [Hello Go](#hello-go)
  - [Main function](#main-function)
  - [Keywords](#keywords)
  - [Built-in functions](#built-in-functions)
  - [Basic types](#basic-types)
  - [Zero values](#zero-values)
  - [Names](#names)
  - [Constants](#constants)
  - [Variables](#variables)
  - [New and make](#new-and-make)
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
  - [Methods](#methods)
  - [Testing](#testing)
  - [Defer](#defer)
  - [Error handling](#error-handling)
  - [Panic and recover](#panic-and-recover)
  - [Modules](#modules)
  - [Channels](#channels)
  - [Generics](#generics)
  - [Useful links](#useful-links)

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
- [dave.cheney/what-is-the-zero-value-and-why-is-it-useful](https://dave.cheney.net/2013/01/19/what-is-the-zero-value-and-why-is-it-useful)

## Names
1. "Avoid package names like base, common, or util"
1. "A package's name should describe its purpose"
1. "A variable's name should describe its content"
1. "Use the smallest scope possible, declare variable close to their use"
1. "Functions should be named for the result they return"
1. "Methods should be named for the action they perform"
1. "Methods mutate state, functions transform data"
1. "The visibility of a name outside a package is determined by whether its first character is upper case"
1. "Finally, the convention in Go is to use MixedCaps or mixedCaps rather than underscores to write multiword names"

- [doc/effective_go#names](https://go.dev/doc/effective_go#names)
- [blog/package-names](https://go.dev/blog/package-names)
- [talks/names](https://talks.golang.org/2014/names.slide)
- [dave.cheney/practical-go](https://dave.cheney.net/practical-go/presentations/qcon-china.html)
- [dave.cheney/avoid-package-names-like-base-util-or-common](https://dave.cheney.net/2019/01/08/avoid-package-names-like-base-util-or-common)
- [dave.cheney/you-shouldnt-name-your-variables-after-their-types-for-the-same-reason-you-wouldnt-name-your-pets-dog-or-cat](https://dave.cheney.net/2019/01/29/you-shouldnt-name-your-variables-after-their-types-for-the-same-reason-you-wouldnt-name-your-pets-dog-or-cat)
- [yt/What's in a name?, Dave Cheney](https://www.youtube.com/watch?v=GXijarXgHHY)
- [wiki/CodeReviewComments#initialisms](https://go.dev/wiki/CodeReviewComments#initialisms)

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

## New and make
```
The new built-in function allocates memory. The first argument is a type,
not a value, and the value returned is a pointer to a newly allocated zero
value of that type.
```
- [ref/spec#Allocation](https://go.dev/ref/spec#Allocation)
- [doc/effective_go#allocation_new](https://go.dev/doc/effective_go#allocation_new)
```go
package main

func main() {
	p := new(int) // pointer to int
	println(p)
	println(*p)
}
```

```
The make built-in function allocates and initializes an object of type
slice, map, or chan (only). Like new, the first argument is a type, not
a value. Unlike new, make's return type is the same as the type of its
argument, not a pointer to it.
```
- [ref/spec#Making_slices_maps_and_channels](https://go.dev/ref/spec#Making_slices_maps_and_channels)
- [doc/effective_go#allocation_make](https://go.dev/doc/effective_go#allocation_make)
- [tour/moretypes/13](https://go.dev/tour/moretypes/13)
- [dave.cheney/go-has-both-make-and-new-functions-what-gives](https://dave.cheney.net/2014/08/17/go-has-both-make-and-new-functions-what-gives)
```go
package main

func main() {
	s := make([]int, 10) // slice with len(s) == cap(s) == 10
	println(s)
	println(s[0])
}
```

## Packages
- [ref/spec#Packages](https://go.dev/ref/spec#Packages)
- [tour/basics/1](https://go.dev/tour/basics/1)
- [doc/effective_go#package-names](https://go.dev/doc/effective_go#package-names)
- [dave.cheney/avoid-package-names-like-base-util-or-common](https://dave.cheney.net/2019/01/08/avoid-package-names-like-base-util-or-common)
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
- [dave.cheney/the-empty-struct](https://dave.cheney.net/2014/03/25/the-empty-struct)

## Methods
```go
package main

import "fmt"

type rectangle struct {
	description string
	a, b        int
}

func (r rectangle) area() int {
	return r.a * r.b
}

func main() {
	r := rectangle{"square", 2, 2}
	fmt.Println(r.area())
}
```
- [ref/spec#](https://go.dev/ref/spec#Method_declarations)
- [ref/spec#](https://go.dev/ref/spec#Method_expressions)
- [ref/spec#](https://go.dev/ref/spec#Method_values)
- [tour/methods/1](https://go.dev/tour/methods/1)
- [tour/methods/2](https://go.dev/tour/methods/2)
- [tour/methods/3](https://go.dev/tour/methods/3)
- [doc/effective_go#methods](https://go.dev/doc/effective_go#methods)
- [doc/faq#Functions_methods](https://go.dev/doc/faq#Functions_methods)
- [doc/faq#methods_on_basics](https://go.dev/doc/faq#methods_on_basics)
- [doc/faq#overloading](https://go.dev/doc/faq#overloading)

## Testing
```go
package main

import (
	"testing"
)

func TestArea(t *testing.T) {
	tests := []struct {
		name string
		rect rectangle
		want int
	}{
		{
			"empty rectangle",
			rectangle{"empty", 0, 1},
			0,
		},
		{
			"special rectangle",
			rectangle{"square", 2, 2},
			4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.rect.area(); got != tt.want {
				t.Errorf("%s: want[%v] != got[%v]", tt.name, tt.want, got)
			}
		})
	}
}
```
- [doc/faq#Packages_Testing](https://go.dev/doc/faq#Packages_Testing)
- [doc/faq#How_do_I_write_a_unit_test](https://go.dev/doc/faq#How_do_I_write_a_unit_test)
- [doc/faq#testing_framework](https://go.dev/doc/faq#testing_framework)
- [doc/code#Testing](https://go.dev/doc/code#Testing)
- [doc/faq#assertions](https://go.dev/doc/faq#assertions)
- [wiki/TableDrivenTests](https://go.dev/wiki/TableDrivenTests)
- [wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables](https://go.dev/wiki/CommonMistakes#using-goroutines-on-loop-iterator-variables)
- [dave.cheney/prefer-table-driven-tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
- [dave.cheney/test-fixtures-in-go](https://dave.cheney.net/2016/05/10/test-fixtures-in-go)

"Files whose names begin with `_` (including `_test.go`) or `.` are ignored."
- [pkg/cmd/go#hdr-Test_packages](https://pkg.go.dev/cmd/go#hdr-Test_packages)

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

## Modules
```sh
go mod init
go mod tidy
```
- [ref/mod](https://go.dev/ref/mod)
- [blog/using-go-modules](https://go.dev/blog/using-go-modules)
- [blog/module-changes](https://go.dev/blog/go116-module-changes)

## Channels
```go
package main

func main() {
	ch := make(chan string)
	go func(chan string) {
		ch <- "Hello, Go!" // write to channel
	}(ch)
	println(<-ch) // read from channel
}
```
- [ref/spec#Channel_types](https://go.dev/ref/spec#Channel_types)
- [ref/spec#Making_slices_maps_and_channels](https://go.dev/ref/spec#Making_slices_maps_and_channels)
- [doc/effective_go#channels)](https://go.dev/doc/effective_go#channels)
- [doc/effective_go#chan_of_chan](https://go.dev/doc/effective_go#chan_of_chan)

## Generics
```go
package main

func add[T int | string](s []T) T {
	var r T

	for _, v := range s {
		r += v
	}

	return r
}

func main() {
	i := []int{1, 2, 3}
	s := []string{"One", "Two", "Three"}
	println(add(i))
	println(add(s))
}
```
- [An Introduction To Generics](https://go.dev/blog/intro-generics)
- [Tutorial: Getting started with generics](https://go.dev/doc/tutorial/generics)
- [When did Go get generic types?](https://go.dev/doc/faq#generics)
- [Why was Go initially released without generic types?](https://go.dev/doc/faq#beginning_generics)
- [How are generics implemented in Go?](https://go.dev/doc/faq#generics_implementation)
- [How do generics in Go compare to generics in other languages?](https://go.dev/doc/faq#generics_comparison)
- [Type Parameters Proposal](https://go.googlesource.com/proposal/+/master/design/43651-type-parameters.md)

### Useful links
- [Release dashboard](https://dev.golang.org/release)
- [Release History](https://go.dev/doc/devel/release)
- [The Go Playground](https://go.dev/play/)
- [A tour of Go](https://go.dev/tour/list)
- [Effective Go](https://go.dev/doc/effective_go)
- [Frequently Asked Questions](https://go.dev/doc/faq)
- [The Go Programming Language Specification](https://go.dev/ref/spec)
- [Installing multiple Go versions](https://go.dev/doc/manage-install#installing-multiple)
- [How to Write Go Code](https://go.dev/doc/code)
- [Optional environment variables](https://go.dev/doc/install/source#environment)
- [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
- [Standard library](https://pkg.go.dev/std)
- [Go command](https://pkg.go.dev/cmd/go)
