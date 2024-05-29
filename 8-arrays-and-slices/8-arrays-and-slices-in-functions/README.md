# Arrays and Slices in Functions

All of this array and slice functionality is useful, but we need to be able to use it outside of `main`.

We can pass arrays or slices into functions as parameters. To pass an array parameter into a function, we provide a local name, square brackets, and the data type. The difference between slice and array parameters is whether the number of elements is stated:

```go
func printFirstLastArray(array [4]int) {
  fmt.Println("First", array[0])
  fmt.Println("Last", array[3])
}

func printFirstLastSlice(slice []int) {
  length := len(slice)
  if (length > 0) {
      fmt.Println("First", slice[0])
      fmt.Println("Last", slice[length-1])
  }
}
```

Due to Go being a pass by value language, modifying a normal array parameter won’t create permanent change. Sometimes this can be useful in performing local calculations.

```go
// Changes to the array will only be local to the function
func changeFirst(array [4]int, value int) {
  array[0] = value
}
```

In order to retain changes, a slice can be used. Find out why this happens for slices and not array

```go
// Changes to the slice parameter will be permanent
func changeFirst(slice []int, value int) {
  if (len(slice) > 0) {
      slice[0] = value
  }
}
```
