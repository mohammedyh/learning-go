# Pointers to a Struct

Without pointers, when a variable is passed into a function, only a copy of it is used inside the function. We can use pointers to modify values in our structs within a function.

We will use an Employee `struct` as an example. We must first create an instance of `Employee` and then we create a pointer that will point to this instance:

```go
type Employee struct {
  firstName string
  lastName string
  age int
  title string
}

func main() {
  steve := Employee{“Steve”, “Stevens”, 34, “Junior Manager”}
  pointerToSteve := &steve
}
```

We can now use this pointer to change the values of the fields for steve. There are two ways to do this in Go:

```go
(*pointerToSteve).firstName
```

Or a simpler, recommended method:

```go
pointerToSteve.firstName
```
> Accessing the fields of a pointer to a struct does not require dereferencing

We can use these pointers to modify structs in our functions. Consider the following example:

```go
// Method is only accessible on pointer(s) Rectangle
func (rectangle *Rectangle) modify(newLength float32) {
  rectangle.length = newLength
}
```

Notice that inside the function `modify()` that `rectangle` is also a pointer. It is dereferenced without the use of the dereferencing operator just like `pointerToSteve`!
