# app

This is a test app, written in [Golang](https://golang.org/).

### Endpoints

There are 2 endpoints exposed in the app;

```
GET /health-check
```

Used to evaluate the health of the app. This returns a 204 if the app is healthy.

And

```
POST /test

Example request:
{
    "field_1": "",
    "field_2": "",
    "field_3": ""
}
```

Used to submit some test data. All fields are optional, and a 200 is returned if the data is processed correctly.

### Environment variables

```
PORT=8123        # the port on which to run
LOG_LEVEL=info   # possible values are 'trace', 'debug', 'info', 'warn' or 'error'
```

### Compilation and testing

This app contains a `Makefile` for convenience.

To 'clean' the app, run `make clean` at the root of the project.

To build the app, run `make build` at the root of the project.

To run unit tests, run `make test-unit` at the root of the project.

To run multiple commands, simply chain them like so: `make clean build test-unit`.

##### Docker

Alternatively, to bake a Docker image, run `docker build -t app .`.