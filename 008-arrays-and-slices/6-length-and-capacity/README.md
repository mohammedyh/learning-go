# Length and Capacity

## Length

The `len()` function returns the length of the provided array or slice.

Arrays only have a length, but slices have a length and a **capacity**.

## Capacity

A slice is resizeable, so there is a difference between:

- Its length - the current number of elements it holds
- Its capacity - the number of elements it _can_ hold before needing to resize itself

A slice’s capacity can be accessed through the `cap` function

```go
slice := []string{"Fido", "Fifi", "FruFru"}
// Slice begins at length 3 and capacity 3
fmt.Println(slice, len(slice), cap(slice))
// Output: [Fido Fifi FruFru] 3 3
slice = append(slice, "FroFro")
// After appending an element when the slice is at capacity
// The slice will double in capacity, but increase its length by 1
fmt.Println(slice, len(slice), cap(slice))
// Output: [Fido Fifi FruFru FroFro] 4 6
```

Note how in the above example, when an element was added to a slice which was at full capacity the following occured:

- The new element was still able to be added
- The length increased to fit the new element (length + 1)
- The capacity doubled in size

All of this happens automatically using slices, while this is not possible with arrays!
