# Multiple Return Values

In Go, functions are able to return multiple values.

To do this, first specify multiple return types, wrap the types in parentheses, where you would normally place a single return type.

Then, we must return two values that match the return types, by writing, for example, `return valueOne, valueTwo`

```go
func GPA(midtermGrade float32, finalGrade float32) (string, float32) {
  averageGrade := (midtermGrade + finalGrade) / 2
  var gradeLetter string
  if averageGrade > 90 {
    gradeLetter = "A"
  } else if averageGrade > 80 {
    gradeLetter = "B"
  } else if averageGrade > 70 {
    gradeLetter = "C"
  } else if averageGrade > 60 {
    gradeLetter = "D"
  } else {
    gradeLetter = "F"
  }

  return gradeLetter, averageGrade
}
```
