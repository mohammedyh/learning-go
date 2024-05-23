# Randomizing

Previously, hard coded values (values that don’t change) were used for variables before the conditionals that checked on these values. For example:

```go
alarmRinging := true
if alarmRinging {
    fmt.Println("Turn off the alarm!!")
}
```

We knew the condition would be true and the print statement would execute. This level of certainty is usually NOT the reason why we would use a conditional. Instead, we create conditionals to account for different conditions and different possible outcomes.

So let’s introduce some uncertainty by generating a random number. Go has a `math/rand` library that helps us generate a random integer.

using rand and the `Intn()` function to print out a random number. With the argument of 100, the maximum value that the method will return is 99. Looking at the entire line `fmt.Println(rand.Intn(100))`, it should print a random integer from 0 to 99. However, if we run our program multiple times, we’ll find that it always prints 81.
