# Importing Multiple Packages

To import multiple packages, we can add multiple import statements:

```go
import "package1"
import "package2"
```

Or use a single import with parentheses:

```go
import (
  "fmt"
  "os"
)
```

# Package Aliases

We can also provide an alias to a package by specifying an alias name before the package name.

```go
import (
  "fmt"
  t "time"
)

func main() {
  fmt.Println(t.Now())
}
```

In the example above we’ve aliased `time` as `t` and called the `Now` function using the alias `t`.
