# Hello Go

A small Go learning repository containing a basic HTTP API and introductory Go examples.

## Requirements

- Go 1.27.1 for the HTTP API (`tutorial/go.mod`)
- Go 1.26.4 for the BornToDev examples (`go_basic_borntodev/go.mod`)

Check the installed version:

```sh
go version
```

## Repository structure

```text
hello-go/
├── README.md
├── tutorial/
│   ├── go.mod
│   └── main.go
├── go_basic_kongruksiam/
│   ├── 1_go_variables.go
│   ├── 2_go_constants.go
│   ├── 3_go_arrays.go
│   ├── 4_go_slices.go
│   ├── 5_go_maps.go
│   ├── 6_go_loop.go
│   ├── 7_go_functions.go
│   ├── 8_go_struct.go
│   ├── 9_go_package.go
│   └── calculator/app.go
└── go_basic_borntodev/
    ├── go.mod
    ├── main.go
    └── 01_go_arrays.go
```

## Run the HTTP API

The API module is located in `tutorial/`:

```sh
cd tutorial
go run .
```

The server listens on `http://localhost:8080`. Test it from another terminal:

```sh
curl -i http://localhost:8080/hello
```

`GET /hello` returns `200 OK` with a greeting. Other HTTP methods return
`405 Method Not Allowed`.

### Verify the API

Run these commands from `tutorial/`:

```sh
go test .
go vet .
go build -o hello-go .
```

## Run the basic examples

The files in `go_basic_kongruksiam/` are separate `package main` examples.
Run one file at a time from the repository root:

```sh
go run go_basic_kongruksiam/1_go_variables.go
go run go_basic_kongruksiam/3_go_arrays.go
go run go_basic_kongruksiam/8_go_struct.go
```

| File | Topic |
|---|---|
| `1_go_variables.go` | Variables and data types |
| `2_go_constants.go` | Constants |
| `3_go_arrays.go` | Arrays |
| `4_go_slices.go` | Slices, `len`, `cap`, and `append` |
| `5_go_maps.go` | Maps |
| `6_go_loop.go` | `for` and `range` |
| `7_go_functions.go` | Functions and return values |
| `8_go_struct.go` | Structs and slices of structs |
| `9_go_package.go` | Calling the `calculator` package |

Do not run all tutorial files together because each file defines its own
`main` function.

The examples are based on these learning resources:

- [Kongruksiam Go tutorial](https://youtu.be/pytqhPDTjnQ?si=ny6qhAodjOly_mJG)
- [BornToDev Go tutorial](https://youtu.be/fjEB75Xotxc?si=pJxwbRGYnWqfOg5n)

## Notes

- The repository contains multiple independent Go modules. Run commands from
  the module directory required by that example.
- `9_go_package.go` currently uses the historical import path
  `hello-go/go_basic/calculator`; update the import path before running it
  after the directory rename to `go_basic_kongruksiam`.
- This is a learning project and is not production-ready. A production API
  should add HTTPS, rate limiting, server timeouts, structured logging, and
  monitoring.
