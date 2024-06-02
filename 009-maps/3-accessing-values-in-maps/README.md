# Accessing Values in Maps

Maps make it easy to look up values and store the value in a variable for further use:

```go
variable := yourMap[keyValue]
```

If a key is not in the map, a default value for value type is returned. We can also get a second return value to determine if the key is in the map.

```go
customer,status := customers["billy"]

if status {
  fmt.Println("we found the customer")
} else {
  fmt.Println("no such customer!")
}
```
