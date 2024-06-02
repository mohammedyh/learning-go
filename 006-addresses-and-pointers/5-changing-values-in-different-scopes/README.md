# Changing Values in Different Scopes

Using our knowledge of addresses, pointers, and dereferencing, let’s return to the [initial problem](../1-the-point-of-addresses-and-pointers/README.md) How can we change the value of a variable when we’re in a different scope?

Code for reference

```go
func addHundred(num int) {
  num += 100
}

func main() {
  x := 1
  addHundred(x)
  fmt.Println(x)
  // Output: 1
}
```

Even though we call `addHundred(x)`, the value of `x` doesn’t change! Why is that?

Remember, Go is a pass-by-value language. When we call `addHundred(x)` we’re providing `addHundred()` with a value of `1`. We’re not actually providing the address of `x` for `addHundred()` to go in and change the value stored there.

In essence this is what happens:

```go
func addHundred(num int) {
  // 4. addHundred receives 1 as the argument, and does 1 += 100
  // but as this isn't assigned back to anything, nothing happens
  num += 100
}

// 1. main() is the entry point so program starts here
func main() {
  // 2. create a variable x, of type int, set to 1
  x := 1
  // 3. call addHundred with the *value* of x, which is 1
  // to Go, it looks like addHundred(1)
  addHundred(x)
  // 5. Print x, which equals 1
  fmt.Println(x)
  // Output: 1
}
```

To change the value of `x` using a function, we need to do the following:

```go
func addHundred(num *int) {
	*num += 100
}

func main() {
	x := 1
	addHundred(&x)

	fmt.Println(x)
}
```

`addHundred()` takes an `int` pointer, meaning, when `addHundred()` is called, it expects a memory address as an argument (which is a pointer to an `int` value).

Inside the function, we are _dereferencing_ the pointer so we can assign a new value using the `*` symbol. If we were to add `fmt.Println(num)`, we would see a reference to the memory address of `num`, that's why we need to use the `*` to **de**reference it.

Since `addHundred()` expects a pointer (and pointers are variables that store an address) we need to provide `addHundred()` with the address of `x` using `&`.
