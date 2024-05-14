# Packages

Now that we understand how to compile and run Go programs, let’s take a look at Go packages.

Projects can contain many `.go` files, organized into packages. Each package is like a directory: `.go` files to do with one part of your program can go in one package, other `.go` files to do with something else can go into another package.

Example: If we were writing a calculator program, we could put the files for the calculation in a package called `calc` and the files for input & output in a package called `io`.

## The Program So Far

```go
package main

import "fmt"

func main () {
  fmt.Println("Hello World")
}
```

### Package Declaration

`package main` is a package declaration, it tells the Go compiler which package this `.go` file belongs to. If this package declaration is `package main`, then this program will be compiled into an executable. `package main` is special to Go, it's where Go looks as the entry point for the application / project.

The `import` keyword lets us use code from other packages. Here we are using the `Println` function on the `fmt` package (`fmt` can also be referred to as a 'standard library'). It's important to note that double quotes are used instead of single quotes - the latter will produce an error. More on this later.

If you are curious on where the `fmt` package comes from, you can go to the location of the Go installation, which is usually `/usr/local/go`. Look in the `src` folder, where you will see many subfolders like `fmt`, `os` and more. These are Go's standard libraries. It's a good idea to read these files to see how the creators of Go structure their files, what naming conventions they follow etc. The files have helpful comments throughout making them easier to follow.
