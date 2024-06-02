# The Point of Addresses and Pointers

Go is a pass-by-value language. For example, we're passing functions the value of an argument.

When we’re calling a function with an argument, the Go compiler is strictly using the value of the argument rather than the argument itself. Because of this feature (pass-by-value), the changes that take place in our function, stay within that function.

But, we **do** have the ability to change values from different scopes. To do so, we need to make use of:

- addresses
- pointers
- dereferencing
