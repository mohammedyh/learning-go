# Getting User Input

The `Scan()` function allows us to get user input.

The `Scan()` function scans text, read from stdin, and stores space-separated values in separate arguments. The code in the "single argument" example only saves the first space separated value. To save more, you need to declare those variables like shown in the "multiple arguments" example.

An interesting point about the above code is that, although there are two `Scan()` lines, if you enter 2 (space-separated) values on a single line, and press Enter, `Scan()` understands that as 2 arguments, and then terminates the scanning.

`Scan()` expects memory _addresses_ / _locations_ for arguments, which is what the `&` does; it gets the variable's memory location.

> Note: Doing `fmt.Println(&response)` returns the memory address for the `response` variable. Example: 0x1400008e050
