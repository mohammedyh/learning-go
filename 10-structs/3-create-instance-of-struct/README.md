# Creating an Instance of a Struct

To use a struct we just defined, we have to create an instance of it; which can be seen below:

```go
type Point struct {
  x int
  y int
}

p1 := Point{x: 10, y: 12}
```

Go also allows us to rely on default values as well, we can omit certain fields by labelling the field(s) we want to populate.

```go
p1 := Point{y: 10}
// x will be set to 0
```

We can even omit all fields to rely on default values for the field type.

```go
p1 := Point{}
// x and y will be set to 0
```
