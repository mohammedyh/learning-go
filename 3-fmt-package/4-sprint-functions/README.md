# The `Sprint` and `Sprintln` Functions

`Sprint()` and `Sprintln()` format strings but don't print them out. Instead, they return the formatted string which can be assigned to a variable. Spaces are added between operands / arguments when neither is a string.

`Sprintln()` is like `Sprintf()`, but adds spaces between arguments regardless of their type and appends a newline at the end.

> Note: `Print()`, `Println()` and `Printf()` also return values, the number of bytes written and any write error encountered.
