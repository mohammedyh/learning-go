# Returning Values from Functions

Use the `return` keyword to pass a value to a function's call site - the place where it is called / invoked.

Functions can be given a return type, telling us the type of a value that will be returned by the function. If we want to `return` something, we must specify a return type on the function.

To store the result / return value of a function, it can be assigned to variable.

Any code after the `return` statement isn't executed. The `return` keywords tells the function to terminate it's execution there.

```go
func getLengthOfCentralPark() int32 {
  var lengthInBlocks int32
  lengthInBlocks = 51
  return lengthInBlocks
}

func main() {
  centralParkLength := getLengthOfCentralPark()
  fmt.Println(centralParkLength)
  // Output
  : 51
}
```
