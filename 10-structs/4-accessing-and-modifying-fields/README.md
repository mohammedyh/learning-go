# Accessing and Modifying Struct Fields

Let’s say we have an instance of the Student struct:

```go
type Student struct {
  firstName string
  lastName  string
  age       int
  grade     int
}

john := Student{"John", "Smith", 14, 9}
```

We can access individual fields within struct using the `.` (dot) operator. Example:

```go
fmt.Println(john.firstName)
```

We can change the value of a field with an assignment statement

```go
john.age = 15
```
