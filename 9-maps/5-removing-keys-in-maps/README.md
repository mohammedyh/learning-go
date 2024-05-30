# Removing Maps

Sometimes we have a key that we no longer need in our map. Go allows us to remove elements using the delete function:

```go
delete(yourMap, keyValueToDelete)
// E.g.
delete(contacts, "Gary")
```

If we call the delete function with a key that is not in the map nothing bad happens. Other languages may crash or throw an exception!
