# Introduction to Slices

So far we’ve been using arrays, which have a fixed size. If we were to want to store a different number of elements in our array, we’d have to make a whole new one. However, Go provides us with a useful alternative.

Slices are a data collection type similar to arrays, but they have the ability to change their size.

There are multiple ways to create a slice. We can create a slice from an array, or by itself. Let’s start with creating a slice by itself.

Each of the following creates an empty slice:

```go
var numberSlice []int
stringSlice := []string{}

// Creates a slice with elements
names := []string{"Kathryn", "Martin", "Sasha", "Steven"}
```

While this last slice currently has four elements, we would be able to continue to add elements using built-in functions.

Below shows how we can create a slice based on that array.

> Note: Modifying the slice will still update the original array.

```go
array := [5]int{2, 5, 7, 1, 3}
// This is a slice of the whole array
wholeSlice := array[:]
// This is a slice containing the elements at indices 2, 3, and 4
partialSlice := array[2:5]

fmt.Println(wholeSlice)
// Output: [2 5 7 1 3]
fmt.Println(partialSlice)
// Output: [7 1 3]
```

In slices, elements are accessed and modified the same way as arrays.

```go
var names = []string{"Kathryn", "Martin", "Sasha", "Steven"}
names[3] = "Bishop"

fmt.Println(names[1])
// Output: Martin
fmt.Println(names[3])
// Output: Bishop
```

> After experimenting with slices, I found that the elements in a slice, point to the same memory addresses as the elements in the original array.

A lot more information about slices here: https://go.dev/blog/slices
