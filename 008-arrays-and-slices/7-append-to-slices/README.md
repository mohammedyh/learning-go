# Append to Slices

A big aspect of working with slices are that we can add elements and the slice will resize automatically. But how do we add to a slice?

Go provides us with a function, `append` that handles all of the logic of adding to (the end of a slice) and resizing a slice.

```go
books := []string{"Tom Sawyer", "Of Mice and Men"}
books = append(books, "Frankenstein")
books = append(books, "Dracula")
fmt.Println(books)
// Output: [Tom Sawyer Of Mice and Men Frankenstein Dracula]
```

Being able to add new elements to a slice is very important, especially when data comes in over time that we cannot predict.
