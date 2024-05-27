# Dereferencing

We know addresses are where values are stored, and pointers allow us to keep track of addresses. But what if we want the address to store a different value? We can actually use our pointer to access the address and change its value! This action is called _dereferencing_ or _indirecting_.

We’ll need to use the `*` operator again on a pointer and then assign a new value like so:

```go
lyrics := "Moments so dear"
pointerForLyrics := &lyrics

*pointerForLyrics = "Journeys to plan"

fmt.Println(lyrics)
// Output: Journeys to plan
```

The `*` operator on `pointerForLyrics` to dereference it and assign a new value of "Journeys to plan". When we check the value of `lyrics` it’s now "Journeys to plan"!
