# Creating Arrays with Elements

We can use curly braces `{}` to populate the array with elements. Without this, Go will provide default / zero values instead. Remembering the [lesson on default values](../../2-variables-and-types/8-default-values/README.md), the default value for `int`s is 0.

This is how we create arrays with populated elements.

```go
triangleSides := [3]int{15, 26, 30}
triangleAngles := [...]int{30, 60, 90}
```

When creating an array with values, we can have the compiler determine the length automatically using `...` ellipsis syntax; as seen on the second line.

```go
triangleSides := [3]int{15, 26}
fmt.Println(triangleSides)
// Output: [15, 26, 0]
```

If you want to, you can define some values and the unpopulated / unset elements will be given a default value.
