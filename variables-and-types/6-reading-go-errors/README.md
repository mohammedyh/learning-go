# Reading Go Errors

When the Go compiler raises an error the program’s binary cannot get created. Go tries to tell you what the issue is by offering you the following pieces of information:

- name of the file where something went wrong
- line number where Go noticed an issue
- column number on that line where the error occurred.

One common error occurs when the Go compiler sees a defined variable that isn’t used in the program. See `main.go` file.

Here is the output from running `go build main.go`:

\# command-line-arguments
./main.go:3:8: "fmt" imported and not used
./main.go:6:6: numberWheels declared and not used
