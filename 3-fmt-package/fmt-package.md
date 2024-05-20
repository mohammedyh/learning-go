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

The `Scan()` function scans text, read from stdin, and stores space-separated values in separate arguments. The above code only saves the first space separated value. To save more, you need to declare those variables like shown below:

```go
fmt.Println("How are you doing?")

var response1, response2 string

fmt.Scan(&response1)
fmt.Scan(&response2)

fmt.Printf("I'm %v %v, thanks for asking\n", response1, response2)

// Output: How are you doing?
// Input: very well
// Output: I'm very well, thanks for asking
```

An interesting point about the above code is that, although there are two `Scan()` lines, if you enter 2 (space-separated) values on a single line, and press Enter, `Scan()` understands that as 2 arguments, and then terminates the scanning.

If you want to accept multiple arguments, it's more common to write it like so:

```go
fmt.Scan(&response1, &response2)
```

`Scan()` expects memory _addresses_ / _locations_ for arguments, which is what the `&` does.

> Note: Doing `fmt.Println(&response)` returns the memory address for the `response` variable. Example: 0x1400008e050

`Print()` and `Sprint()`:
- Uses default formatting
- Puts spaces between arguments that **are not** strings
- Doesn't append a newline

`Printf()` and `Sprintf()`
- Accepts a format specifier string as first argument.
- Doesn't append a newline

`Println()` and `Sprintln()`
- Uses default formatting
- Put spaces between all arguments
- Does append a newline

`Print` functions return the number of bytes written and any write error.

`Sprint` functions returns the resulting, formatted string.
