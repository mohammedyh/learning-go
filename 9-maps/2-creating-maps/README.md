# Creating Maps

In Go, there are a few ways to create a map.

## Creating a (empty) map with `make`

```go
mapVariableName := make(map[keyType]valueType)
```

## Creating a (empty) map with `var`

```go
var emptyMap map[string]string
```


Creating empty maps is useful when we don’t know what the content of our map will be. But sometimes the content of the map is known ahead of time.

## Creating a map with values

```go
variableName := map[keyType]valueType{
  name1: value1,
  name2: value2,
  name3: value3,
}
```
