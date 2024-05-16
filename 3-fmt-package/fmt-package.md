# `fmt` Package

## Lesson 1 - The `Print` Method

`fmt.Println()` puts a spaces between all arguments and prints an ending newline.
`fmt.Print()` doesn't put spaces between arguments and doesn't add an ending newline; these need to be added manually.

```go
fmt.Print("The answer is", ": ")
fmt.Print("12")

// Output: The answer is: 12
```

## Lesson 2 - The `Printf` Method

`fmt.Printf()` allows us to _interpolate_ strings using _Go verbs_.

```go
animal1, animal2 := "Cat", "Dog"
fmt.Printf("Are you a %v or a %v person?\n", animal1, animal2)

// Output: Are you a Cat or a Dog person

// The equivalent using Println would look like:
fmt.Println("Are you a", animal1, "or a", animal2, "person?")
```

`%v` is the placeholder also known as a _Go verb_ and `%v` specifically returns the value of the supplied arguments

In `Printf()` the position of the arguments, are intuitive, but they do matter.

## Lesson 3 - Different Verbs

- `%T` prints out the type of the second argument
- `%d` lets us interpolate a number into a string. (but `%v` also works, find out why)
- `%f` lets us interpolate floats into a string.
- `%.xf` lets us change how precise the float is, where `x` is a number. Example: `%.2f`

```go
floatExample := 1.75
fmt.Printf("Working with a %T", floatExample)

// ----

yearsOfExp := 3
reqYearsExp := 15
fmt.Printf("I have %d years of Go experience and this job is asking for %d years", yearsOfExp, reqYearsExp)

// ----

stockPrice := 3.50
fmt.Printf("Each share of Gopher feed is $%.2f!", stockPrice)
```

In the `Printf()` function, if an argument is passed but isn't used, the following output is shown:

```go
floatExample := 1.75
fmt.Printf("Working with a ____", floatExample)

// Output: Working with a ____%!(EXTRA float64=1.75).
```

## Lesson 4 - `Sprint` and `Sprintln`

`Sprint()` and `Sprintln()` format strings but don't print them out. Instead, they return the formatted string which can be assigned to a variable. Spaces are added between operands/arguments when neither is a string.

`Sprintln()` is like `Sprintf()`, but adds spaces between arguments regardless of their type and appends a newline at the end.

> Note: `Print()`, `Println()` and `Printf()` also return values, but they return the number of bytes written and any write error encountered.

```go
grade := "100"
compliment := "Great job!"
teacherSays := fmt.Sprint("You scored a ", grade, " on the test! ", compliment)

fmt.Println(teacherSays)

// Output: You scored a 100 on the test! Great job!
```

## Lesson 5 - The `Sprintf` Method

`Sprintf()` allows us to interpolate a string without printing it.

- Accepts a format string (using Go verbs like %v, %d, etc), and all necessary arguments afterwards
- Returns the formatted string

```go
correctAns := "A"
answer := fmt.Sprintf("And the correct answer is... %v!", correctAns)

fmt.Println(answer)

// Output: And the correct answer is... A!
```

## Lesson 6 - Getting User Input

The `Scan()` method allows us to get user input, below is an example.

```go
fmt.Println("How are you doing?")

var response string
fmt.Scan(&response)

fmt.Printf("I'm %v.", response)
```

First, "How are you doing?" is printed to the terminal. Then we declare a variable, `response` with the type `string`. `fmt.Scan(&response)` takes the first value before a space and stores it in response. In the terminal, we would see:

`How are you doing?`

The `Scan()` function scans text and stores space separated values. The above code snippet would only save the first space separated value. If you need to save more, you need to declare those variables.

`Scan()` expects _addresses_ / _memory locations_ for arguments, which is what the `&` does.

> Note: Doing `fmt.Println(&response)`
