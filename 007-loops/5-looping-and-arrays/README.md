# Looping and Arrays

> Note: For some reason the course / platform decided to show the syntax for arrays, slices and maps before actually covering them. Not a big deal but worth mentioning that arrays, slices and maps are covered the proceeding lessons.

Quick reference:

- Arrays: fixed-length sequences of elements of the same data type. E.g. `var nums = [5]int{1, 2, 3, 4, 5}`
- Slices: dynamically-sized, flexible views into arrays. _A slice is not an array. A slice describes a piece of an array. (from Go docs)_. E.g. `var names = []string{"Elliot", "Zach", "Michael", "Theo"}`
- Maps: data structure that store unordered collections of key-value pairs. E.g. `var orders = map[string]string{
  "John": "Cheeseburgers",
  "Janet": "Hamburgers",
  "Jordan": "Fries",
}`

Each map and array has a set amount of items that they contain. In Go, the `range` keyword can be used to work through these items one at a time within a loop.

```go
letters := []string{"A", "B", "C", "D"}
for index, value := range letters {
  fmt.Println("Index:", index, "Value:", value)
}
```

Would output the following:

```
Index: 0 Value: A
Index: 1 Value: B
Index: 2 Value: C
Index: 3 Value: D
```

The `range` keyword is used here similarly to the initial statement of a definite loop. It lets the programmer assign two new variables for the index and value of each item in the array.

The behavior is the same for maps. But as they don’t have an index, range provides the key and value pairs for each item instead.
