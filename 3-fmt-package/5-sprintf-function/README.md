# The `Sprintf` Method

`Sprintf()` allows us to interpolate a string like `Printf()`, but without printing it.

- It accepts a format string (using Go verbs like %v, %d, etc), and all necessary arguments afterwards
- It returns the formatted string

Here are differences between the `Print()` and `Sprint()` function variations

`Print()` and `Sprint()`:
- Uses default formatting for operands
- Puts spaces between arguments that **are not** strings
- Doesn't append a newline

`Printf()` and `Sprintf()`
- Accepts a format specifier string as first argument
- Doesn't append a newline

`Println()` and `Sprintln()`
- Uses default formatting for operands
- Puts spaces between all arguments
- Does append a newline

`Print` functions return the number of bytes written and any write error.

`Sprint` functions returns the formatted string.
