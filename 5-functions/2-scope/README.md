# Scope

Scope is a concept that refers to where the values and functions are defined and where they can be accessed. For instance, when a variable is defined within a function, that variable is only accessible within that function. Trying to access the same variable from a different function, results in an error. Each function has its own specific scope.

```go
package main

import "fmt"

// Global scope

func performAddition() {
  // performAddition local scope
  x := 5
  y := 7
  fmt.Println("The sum of", x, "and", y, "is", x + y)
}

func main() {
  // main local scope
  performAddition()
  fmt.Println("What if", x, "was different?")
  // Output: ./main.go: undefined: x
}
```
