# Hello Go

This project contains a sample HTTP API and basic Go tutorials in `go_basic/`.

## Requirements

- Go **1.27.1**, as specified in `tutorial/go.mod`
- macOS or another operating system supported by Go

Check the installed version:

```sh
go version
```

## Project structure

```text
hello-go/
├── README.md
├── tutorial/
│   ├── go.mod
│   └── main.go
└── go_basic/
    ├── 1.go_variables.go
    ├── 2.go_constants.go
    ├── 3.go_arrays.go
    ├── 4.go_slices.go
    ├── 5.go_maps.go
    ├── 6.go_loop.go
    ├── 7.go_functions.go
    ├── 8.go_struct.go
    ├── 9.go_package.go
    └── calculator/app.go
```

## Run the HTTP API

Change to the `tutorial/` directory first because that is where `go.mod` is located:

```sh
cd tutorial
go run .
```

The server listens on `http://localhost:8080`. Press `Ctrl+C` to stop it. From another terminal, call the endpoint with:

```sh
curl -i http://localhost:8080/hello
```

`GET /hello` returns `200 OK` with the following message:

```text
Hello from Go on Omarchy!
```

Other methods sent to `/hello` return `405 Method Not Allowed`:

```sh
curl -i -X POST http://localhost:8080/hello
```

## Build and verify the API

Run these commands from the `tutorial/` directory:

```sh
go test .
go vet .
go build -o hello-go .
./hello-go
```

## Basic Go tutorials

Each file has its own `package main` and `func main()`. Run files individually from the project root:

```sh
go run go_basic/1.go_variables.go
go run go_basic/3.go_arrays.go
go run go_basic/8.go_struct.go
```

| File | Topic |
|---|---|
| `1.go_variables.go` | Variables and data types |
| `2.go_constants.go` | Constants |
| `3.go_arrays.go` | Arrays |
| `4.go_slices.go` | Slices, `len`, `cap`, and `append` |
| `5.go_maps.go` | Key/value maps |
| `6.go_loop.go` | `for` and `range` |
| `7.go_functions.go` | Functions and return values |
| `8.go_struct.go` | Structs and a `Person` slice |
| `9.go_package.go` | Calling the `calculator` package |

### Note about example 9

This example imports `hello-go/go_basic/calculator`, but the current module is in `tutorial/`. Create a module for the tutorials first:

```sh
cd go_basic
go mod init hello-go/go_basic
go run 9.go_package.go
```

Expected output:

```text
7
5
```

The `go mod init` command creates `go_basic/go.mod`. To undo this setup, remove that file manually after confirming that it is no longer needed.

## Troubleshooting

- `go.mod file not found`: Change to `tutorial/` before running the API.
- `main redeclared`: Do not run all files in `go_basic/` together; specify one file at a time.
- `package ... is not in std` in example 9: Create `go_basic/go.mod` using the command above.
- `address already in use`: Stop the existing server with `Ctrl+C` or change the port in `tutorial/main.go`.

## Scope

This code is intended for learning and uses only the Go standard library. The API is not production-ready; production deployment should add HTTPS, rate limiting, server timeouts, logging, and monitoring.

## Learning resource

The basic examples reference this [Go tutorial video](https://youtu.be/pytqhPDTjnQ), as listed in `go_basic/readme.md`.
