# Using Functions

Functions are defined using the `func` keyword, followed by the function name.

> You cannot define a named function inside the `main` function, only outside of it.

Functions are _declared_ at "top level" package scope. If you want to define a function inside a function (like `main`), assign to a variable. This is a way to achieve some of the benefits of nested functions while maintaining the language's simplicity and avoiding potential issues. This is known as an anonymous function since it doesn't have a name after the `func` keyword.

```go
func main() {
  addNums := func(x, y int) int {
    return x + y
  }

  fmt.Println(addNums(2, 2))
  // Output: 4
}
```

This design decision was made by the Go language creators to keep the language simple and avoid potential issues that can arise from nested functions, such as complex scoping rules, potential name conflicts etc.
