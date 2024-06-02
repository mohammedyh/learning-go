# Nested / Embedded Structs

When we have complex groups of fields in our structs, they can be combined into their own struct.

For example, in the `Employee` struct, we have two separate fields for the first and last name. Those two fields can be combined into their own struct called which we'll call `Name`.

```go
type Name struct {
  firstName string
  lastName  string
}

type Employee struct {
  // This is typically referred to as "struct embedding"
  name  Name
  age   int
  title string
}
```

The `Employee` instance would look like this:

```go
person := Employee{
  Name{"Carl", "Carlson"},
  32,
  "Engineer",
}
```

Accessing fields on the `Name` struct, intuitively, is done by chaining together the field names.

```go
fmt.Println(person.name.firstName)
// Output: Carl
```

We can also define the employee struct with the `Name` struct anonymously - meaning without an associated field / variable name. This leverages Go's _anonymous field embedding_.

```go
type Employee struct {
  Name
  age int
  title string
}
```

Composing a struct in this way allows us to access the `firstName` and `lastName` fields directly from the `Employee` struct. However, you **can still** access the fields with the full path using the embedded type name.

```go
person := Employee{
  Name{"Carl", "Carlson"},
  32,
  "Engineer"
}

fmt.Println(person.firstName)
// Output: Carl

// Using the name of the embedded type
fmt.Println(person.Name.lastName)
// Output: Carlson
```

An anonymous field is used to field access easier and leads to cleaner code.
