# Seeding

> Note: When trying to use `rand.Seed(x)` the compiler shows the following message, _"As of Go 1.20 there is no reason to call Seed with a random value. [...]"_

The reason why `rand.Intn(1000)` produces the same values every time you use it is due to how Go seeds or chooses a number as a starting point for generating random numbers. By default, the seed value is 1. We can provide a new seed value using the function `rand.Seed()` and passing in an integer.

> (seed | seeded | seeding) refers to initializing the internal state of the pseudorandom number generator algorithm.
 This seed value acts like a starting point that determines the sequence of random numbers produced.

 However, we encounter the same problem if we pass in a constant like 5. Each time we run our program, we’ll always print the same set of numbers. Therefore, each time we run our program, we need a unique number as a seed. One way to do this is to use time. The reason is because it’s always a different time when our program is run.

We can use the `time` library and call `.Now()` with the `.UnixNano()` function chained to it. Resulting in the difference in time (in nanoseconds) since Janurary 1st, 1970 UTC. The number that we get from `time.Now().UnixNano()` is passed as an argument to `rand.Seed()` resulting in a different seed value each time we run our program.

Additional notes:

- Pseudorandom Numbers: Computers can't generate true randomness. Instead, they use mathematical algorithms to produce pseudorandom numbers, which appear random but are actually deterministic (based on a seed).
- Seeding the Generator: The seed value initializes the algorithm's internal state. Changing the seed significantly alters the sequence of numbers produced.
- Default Seed in Go: By default, `math/rand` often uses a fixed seed value, like the current time when the program starts. This ensures consistency in your results if you need to re-run the program to debug or test something.
