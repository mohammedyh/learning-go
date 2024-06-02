# Pointers

Pointers are variables that specifically store memory addresses.

```go
var pointerForInt *int
```

In the example above `pointerForInt` will store the address of a variable that has an `int` data type. The `*` operator signifies that this variable will store an address and the `int` portion means that the address contains an integer value.

With `pointerForInt` initialized, we can assign it value like so:

```go
// Makes a variable with type "pointer int"
// Points to an address containing an int value
var pointerForInt *int
minutes := 525600

// Stores the memory address of minutes in pointerForInt
pointerForInt = &minutes

fmt.Println(pointerForInt)
// Output: 0xc000018038
```

> It's worth noting that a pointer's zero value is `nil`
