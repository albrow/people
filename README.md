People
======

An simple HTTP/JSON API written in Go.

People uses the following libraries:

- [github.com/albrow/forms](https://github.com/albrow/forms)
- [github.com/albrow/negroni-json-recovery](https://github.com/albrow/negroni-json-recovery)
- [github.com/albrow/zoom](https://github.com/albrow/zoom)
- [github.com/codegangsta/negroni](https://github.com/codegangsta/negroni)
- [github.com/gorilla/mux](https://github.com/gorilla/mux)
- [github.com/unrolled/render](https://github.com/unrolled/render)

## Installation

People requires Go version >= 1.5 with `GO15VENDOREXPERIMENT=1`. As of Go 1.6
the vendor experiment is enabled by default.

Run `go get -u github.com/albrow/people`, which will automatically install the 
source code into the correct location at `$GOPATH/src/github.com/albrow/people`.

To start the server, change into the project root directory and run
`go run main.go`. The server runs on port 3000.

### Redis

Since People uses [Zoom](https://github.com/albrow/zoom) as a datastore, it
needs to connect to a Redis database. Follow these
[instructions to install Redis](http://redis.io/topics/quickstart). People will
connect to Redis on the default port, `localhost:6379`, and use database 1.

## Demo

I strongly recommend [httpie](https://github.com/jkbrzt/httpie) for quickly
sending requests to the server from the command line (e.g.,
`http GET :3000/people`). You could also use `curl` or any other HTTP client.

![Demo](https://cloud.githubusercontent.com/assets/800857/13369002/ef7c6278-dca4-11e5-8dba-39d37b27f2f8.gif)

## API Documentation

### GET /people

Returns a list of all people.

### POST /people

Parameters:

- Age: `int`
- Name: `string`

Creates a new person.

### GET /people/{id}

Gets and returns a single person with the given id.

### PATCH /people/{id}

Parameters:

- Age: `int`
- Name: `string`

Update the person with the given id. Parameters are optional and any that are
not provided will remain unchanged.

### DELETE /people/{id}

Removes the person with the given id from the list.
