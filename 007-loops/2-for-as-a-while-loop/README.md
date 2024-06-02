# For as a While Loop

Sometimes it is impossible to know how many iterations of a loop are needed ahead of time.

For example, imagine if a club is throwing an event to promote new member sign-ups. The club has to process each new sign-up as long as more people keep coming!

In situations like these, indefinite loops can be used. These are loops that repeat while a condition remains true. So with the club example, an indefinite loop could run as long as there are new sign-ups to process.

Some other examples could be:

- A car that accelerates while the gas is being pressed.
- A stove that heats up while the knob is turned on.
- A music player that plays songs while the pause button is not pressed.

All of these examples show some sort of process happening while a condition is met. Most programming languages use the `while` keyword to declare indefinite loops.

However, the behavior of definite and indefinite loops is very similar. For this reason, Go simplifies the language by using only the `for` keyword for both types of loops.

The syntax between the two types of loops is pretty similar, but with some crucial differences.

First, any variables to be used inside the loop need to be initialized ahead of time. Second, any update to these variables has to be manually included within the code block. So the previous printing numbers example using a definite loop:

```go
for number := 0; number < 5; number++ {
  fmt.Println(number)
  // Output: 0, 1, 2, 3, 4
}
```

```go
number := 0 // Initialize a variable to be used inside the loop
for number < 5 {
  fmt.Println(number)
  number++ // Update the variable being used
}
```

Both code snippets produce the same output, they are different ways of implementing the same logic.
