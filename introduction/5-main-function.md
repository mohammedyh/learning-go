# `main` Function

```go
func main () {
  fmt.Println("Hello World")
}
```

We use the `func` keyword to declare the Go function `main`:

- The `func` keyword denotes the start of a function declaration.
- `func` is followed by the name of the function: `main`.
- After the name there are a pair of parentheses `()` for function parameters and a set of curly braces `{}` denoting the function body.
- The function code is written inside the set of curly braces.

## Function Code

The code inside a function is indented. The code invokes the `Println` function in the `fmt` package to print the string "Hello World".

## Invoking Functions

Normally when we write functions, you need to write code to invoke them, otherwise they are unused. The `main` function is different if it resides in the `main` package.

When we have a `main` function in the `main` package, this is special to Go. When compiled (using `go build`) an executable is created, and when run (using `go run`), the executable uses the `main` function as the starting point.

```go
package main

import "fmt"

func main() {
  fmt.Println(firstFunc())
}

// return type(s) can be specified after parentheses
func firstFunc() int8 {
  return -128
}
```
