# The `Printf` Function

`fmt.Printf()` allows us to _interpolate_ strings using _Go verbs_.

`%v` is the placeholder also known as a _Go verb_. `%v` specifically returns the value of the supplied arguments in the default format.

In `Printf()` the position of the arguments, are intuitive, but they do matter. E.g. The first argument, after the format string, is interpolated in the position of the first `%v`.
