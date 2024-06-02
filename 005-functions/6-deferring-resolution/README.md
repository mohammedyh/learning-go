# Deferring Resolution

We can delay a function call to the end of the current scope by using the `defer` keyword. `defer` tells Go to run a function, but at the end of the current function. This is useful for logging, file writing, and other utilities.

> Functions prefixed with the `defer` keyword run after the `return` keyword. A deferred function's arguments are evaluated immediately.

Take the following example:

```go
func calculateTaxes(revenue, deductions, credits float64) float64 {
  defer fmt.Println("Taxes Calculated!")
  taxRate := .06143
  fmt.Println("Calculating Taxes")

  if deductions == 0 || credits == 0 {
    return revenue * taxRate
  }

  taxValue := (revenue - (deductions * credits)) * taxRate
  if taxValue >= 0 {
    return taxValue
  } else {
    return 0
  }
}
```

We want to print "Taxes Calculated!", As we have multiple `return` statements,
