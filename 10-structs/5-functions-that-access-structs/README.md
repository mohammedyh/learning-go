# Functions that Access a Struct

We can use functions to capture logic involving our structs and simplify it.

Structs will often have important operations that can be performed on them. For example, with a struct representing a geometric shape, it would be natural to have functions that compute its area and perimeter.

Let’s say we have a struct describing a rectangle. The rectangle struct will contain two fields: the length and the width. We define this struct:

```go
type Rectangle struct {
  length float32
  width  float32
}
```

We can define a function that computes the area of the rectangle.

```go
func (rectangle Rectangle) area() float32 {
  return Rectangle.length * rectangle.width
}
```

The key thing to notice is the line `(rectangle Rectangle)`. This line signals to Go that the `area()` function belongs to the `Rectangle` `struct`. Note that functions associated with a `struct` are written outside of the `struct`.

> `(rectangle Rectangle)` defines the receiver of the method, without it, there wouldn't be a way to reference the `Rectangle` struct's fields within the method.

We can call the `area` function like so:

```go
rect := Rectangle{10, 20}
rect.area()
```

Defining a function in this way will only pass in a **copy** of `rect`, we can't use the function to alter the value of a field. If we want to modify the value of a `struct` field, we have to pass in a pointer to a `struct`.
