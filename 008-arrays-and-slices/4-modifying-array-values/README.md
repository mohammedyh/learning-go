# Modifying Array Values

The syntax for modifying array values is quite simple

```go
array[index] = value
```

Where index is any valid index in the array, value is any expression we want. Let’s say we had an array, and we decide we want the third element to now be 33.

```go
myArray := [4]int{10, 24, 5, 47}
myArray[2] = 33
```

The content of the array would now be `{10, 24, 33, 47}`

We can use this syntax to modify any valid index, between `0` and the length of the array.
