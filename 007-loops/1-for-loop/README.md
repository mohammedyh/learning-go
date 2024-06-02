# Classic For Loop

Loop allow us to check a condition and repeat code a fixed number of times.

While an if statement only has a conditional expression (like `number < 5`), a definite loop has two additional components:

- An initial statement, defining the starting value for a loop variable. `Ex: number := 0`
- A post statement, defining what happens at the end of each iteration. `Ex: number++`

The initial statement, `number := 0`, creates a new variable to be used within the for loop code block.

The conditional expression, `number < 5`, will stop the loop when number reaches the target value of 5.

The post statement, `number++`, increases the value of the number variable by `1` each time the code block completes.

In this example, the amount of times that the number needed to be printed was known beforehand. But what do we do when we don’t know how many times a loop needs to run?
