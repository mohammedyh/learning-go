# Creating Arrays

When we declare a variable in Go, the compiler:

- Finds space in memory for that variable
- Associates the variable with a name

Using arrays makes the compiler’s job a little more complicated. When we declare a single variable, the compiler needs to find enough space for one of that data type. When we declare an array, the compiler is going to have to find enough space for several of a data type. E.g. 5 signed 64 bit integers

To make this process simple, declaring an array in Go requires that we provide the number of elements. Once declared, we cannot change this number without declaring a new array. The compiler finds enough space for the array’s type, multiplied by the number of elements.

We can create arrays with or without an initial set of elements. We use an array without initial elements when the rest of our program will create the array’s content.

```go
var playerScores [4]int
fmt.Println(playerScores)
// Output: [0 0 0 0]
```

This syntax creates an empty array of integer values with space for 4 elements.
