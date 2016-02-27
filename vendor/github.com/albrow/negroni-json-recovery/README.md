JSON Recovery middleware for negroni
-------------------------------------

[Negroni](https://github.com/codegangsta/negroni) is an idiomatic approach to web middleware in Go.

This recovery middleware catches any panics and wraps them up into a json response, including
the line number where the panic occured. Borrows heavily from the default
[recovery middleware in martini](https://github.com/go-martini/martini/blob/master/recovery.go).

See also https://github.com/go-martini/martini/blob/master/LICENSE

Usage
-----

### Installation

Installation is the same as any other package
```
go get github.com/albrow/negroni-json-recovery
```

Make sure you have imported the package:

```go
import "github.com/albrow/negroni-json-recovery"
```

Then add to the middleware stack:

```go
n.Use(recovery.JSONRecovery(true))
```

In production, you should set fullMessages to 'false' to show a generic error message:

```go
n.Use(recovery.JSONRecovery(false))
```

### Full Example

```go
package main

import (
	"github.com/albrow/negroni-json-recovery"
	"github.com/codegangsta/negroni"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		panic("Oh no! Something went wrong in HandleFunc")
	})

	n := negroni.New(negroni.NewLogger())
	n.Use(recovery.JSONRecovery(true))
	n.UseHandler(mux)
	n.Run(":3000")
}
```

With fullMessages set to 'true', the above code will render the following response:

```json
{
    "Code": 500,
    "Short": "internalError",
    "Errors": [
        "dial tcp 127.0.0.1:6379: connection refused"
    ],
    "From": "/Users/alex/programming/go/src/github.com/albrow/testing/negroni-panic/server.go:12"
}
```

or with fullMessages set to 'false', it will render:

```json
{
    "Code": 500,
    "Short": "internalError",
    "Errors": [
        "Something went wrong :("
    ]
}
```

You can also inspect console for a more detailed message and full stack trace.

### Custom Response Formats

You can change the response format by overriding the Formatter function. Formatter is a function
which accepts a variety of parameters related to the error that occurred and returns an empty interface.
The interface will be converted to JSON and rendered in the response.

Here's an example of a custom formatter which only spits out the error message and no other information.

```go
recovery.Formatter = func(errMsg string, stack []byte, file string, line int, fullMessages bool) interface{} {
	return map[string]string{
		"error": errMsg,
	}
}
```

With this custom formatter, the rendered message would look like this:

```json
{
	"error": "dial tcp 127.0.0.1:6379: connection refused"
}
```
