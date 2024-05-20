# Formatting Verbs

- `%T` prints out the type of the second argument
- `%d` lets us interpolate a number into a string.
- `%f` lets us interpolate floats into a string.
- `%.xf` lets us change how precise the float is, where `x` is a number. Example: `%.2f`

`%v` is a more general formatting verb. It works for any value type, including integers, strings, structs, etc. The default format for `%v` is:

```
bool:                    %t - the word true or false
int, int8 etc.:          %d - base 10
uint, uint8 etc.:        %d, %#x if printed with %#v - base 10
float32, complex64, etc: %g - %e for large exponents, %f otherwise
string:                  %s - uninterpreted bytes of the string
```
