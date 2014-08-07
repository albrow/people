JSON Recovery middleware for negroni
Catches any panics and wraps them up into a json response.
Borrows heavily from the default recovery middleware in martini:
https://github.com/go-martini/martini/blob/master/recovery.go

See also https://github.com/go-martini/martini/blob/master/LICENSE