# Function Parameters

We can provide functions with information through parameters. Parameters are variables used within the function (as placeholders) to use in some sort of computation or calculation. When calling a function, we provide it with "arguments", which are the actual values that we want those parameter variables to take. We give our function parameters types when defining the function.

Paramers can be literal values like `5`, `"Earth"` and also variables like `myAge`, `currentPlanet` etc.

```go
func multiplier(x, y int32) int32 {
  return x * y
}
```
