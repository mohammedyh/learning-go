# Defining a Struct

The definition of a `struct` includes its name and its fields. A field is one of the internal variables inside a struct. We use the following template:

```go
// Struct names begin with a capital letter
type Point struct {
  // field and their type go here
  x int
  y int
}
```

> You _can_ make structs using a lowercase first letter, but it goes against the general practice and linters may show warning / errors about it.
