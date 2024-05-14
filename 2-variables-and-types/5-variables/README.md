# Variables

Variables are defined with the `var` keyword and two pieces of information: the name and the type of data stored in the variable. Since variables can be updated we don’t need to initialize it with a value.

The convention is to name variables using camelCase. Things that start with an uppercase like functions and variables mean something special in Go. More on this later.

## Updating Variables

Variables are different from constants because we can update them. This update feature becomes incredibly important when we need to use the original value of a variable for a calculation (or any general manipulation) and then update the variable to store the newly calculated value.

Go has assignment operations that perform calculations and update a variable’s value:

- \+= to add to the variable.
- \-= to subtract from the variable.
- \*= to multiply the variable by a factor.
- /= to divide the variable by a dividend.

## Multiple Variable Declaration

Go allows us to declare multiple variables on a single line, with a few different syntaxes.
