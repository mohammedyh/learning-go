# Break and Continue

The `break` keyword allows us to stop the loop at the current iteration.

We can use the `break` keyword to exit out of the infinite loop in the previous lesson.

```go
var number int

func count() {
	var input int
	fmt.Print("Number of guests: ")
	fmt.Scan(&input)

	number = number + input
	fmt.Println("Total guests:", number)
}

func main() {
	for {
		count()
    break
	}
}
```

The `continue` keyword works similarly, allowing the loop to skip to the next iteration.

```go
const HATED_NUMBER = 3

func main() {
  for number := 0; number < 10; number++ {
    if number == HATED_NUMBER {
      continue
    }
    fmt.Println("I like the number", number)
  }
}
```

The above code prints "I like the number {number}" for every number between 0 and 9, but skips the number 3.

However, using `continue` and `break` statements tend to cause confusion over how a loop will behave. A `break` statement changes when a loop will end. While a `continue` statement changes what will happen in each loop.
