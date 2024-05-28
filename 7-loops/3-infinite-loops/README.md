# Infinite Loops

Sometimes an indefinite loop does not end, creating an infinite loop.

An infinite loop is a loop where the condition can never be reached, causing the loop to run forever.

Sometimes this can be dangerous as it can make a program appear to be frozen to the user. But for something like a web server that constantly needs to be running, these infinite loops are critical.

While an infinite loop can be any loop that does not end, a purposefully created infinite loop is written like this:

```go
for {
  fmt.Println("Infinite loop")
}
// This is never reached
```

There is also a way to stop these loops at a certain point in code
