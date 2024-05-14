# Comments

Comments are useful for:

- Explaining what the code does & why something was done a certain way
- Outlining important or fragile blocks of code, which require extra care
- Noting down what we need to do when writing the code
- Disabling code
- Adding information to be picked up by the `go doc` tool

## Comment Types

### Line Comments

Line comments start with a `//` and the rest of the line is ignored by the compiler.

```go
// This entire line is ignored by the compiler
// fmt.Println("This line does NOT print")
fmt.Println("This gets printed!") // This part gets ignored
```

### Block Comments

Block comments can span multiple lines. They start with a `/*` and end with a `*/`, enveloping everything inside:

```go
/*
This is ignored.
This is also ignored.
fmt.Println("This WON'T print!")
*/
```
