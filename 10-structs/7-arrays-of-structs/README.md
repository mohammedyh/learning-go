# Arrays of Structs

When dealing with many structs of the same type we can use them in an array together. Let’s say we want to create an array of the following points: `{1, 1} {7, 27} {12, 7} {9, 25}`.

We can create an array of `Point` structs like so:

```go
type Point struct {
  x int
  y int
}

points := []Point{{1, 1}, {7, 27}, {12, 7}, {9, 25}}
```

If the points have names, we can also create the array like this:

```go
a = {1, 1}
b = {7, 27}
c = {12, 7}
d = {9, 25}

points := []Point{a, b, c, d}
```

We can access the contents of this array like we would for an ordinary array. We can also access and modify the fields of each of the array elements.

```go
points := []Point{{1, 1}, {7, 27}, {12, 7}, {9, 25}}

fmt.Println(points[0])
// Output: {1, 1}

points := []Point{{1, 1}, {7, 27}, {12, 7}, {9, 25}}

points[1].x = 8
points[1].y = 16

fmt.Println(points[1])
// Output: {8, 16}
```

Arrays of structs allow us access many instances of our structs together in our programs!
