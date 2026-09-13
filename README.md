# Hello Go

A Go learning repository with basic language examples, JSON processing, file I/O, concurrency, and a small HTTP server.

## Requirements

- Go 1.27.1 for the `tutorial` module
- Go 1.26.4 for the `go_basic_borntodev` module
- Go standard library only

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
├── go_basic_borntodev/
│   ├── go.mod
│   ├── 1_array.go
│   ├── 2_slice.go
│   ├── 3_map.go
│   ├── 4_pointer.go
│   ├── 4_pointer_.go
│   ├── 5_struct.go
│   ├── 6_interface.go
│   ├── 7_chanel.go
│   ├── 7_gorutine.go
│   ├── 8_defer.go
│   ├── 9_read.go
│   ├── 10_write.go
│   ├── 11_handle.go
│   ├── 12_jsonMarshal.go
│   ├── 13_jsonUnMarshal.go
│   ├── products.csv
│   └── main.go
└── go_basic_kongruksiam/
    ├── 1_variable.go
    ├── 2_constant.go
    ├── 3_array.go
    ├── 4_slice.go
    ├── 5_map.go
    ├── 6_loop.go
    ├── 7_function.go
    ├── 8_struct.go
    ├── 9_package.go
    └── calculator/app.go
```

## Run the HTTP server

The learning HTTP server is `go_basic_borntodev/11_handle.go`:

```sh
cd go_basic_borntodev
go run 11_handle.go
```

Open the root endpoint in a browser or use `curl` from another terminal:

```sh
curl http://127.0.0.1:8080/
```

Press `Ctrl+C` to stop the server.

The server prints the URL when it starts and reports an error if port `8080`
cannot be opened.

## Run the examples

Each example is a separate `package main` program. Run one file at a time
from the `go_basic_borntodev` directory:

```sh
cd go_basic_borntodev
go run 1_array.go
go run 2_slice.go
go run 3_map.go
go run 4_pointer.go
go run 5_struct.go
go run 6_interface.go
go run 7_chanel.go
go run 7_gorutine.go
go run 8_defer.go
go run 12_jsonMarshal.go
go run 13_jsonUnMarshal.go
```

| File | Topic |
|---|---|
| `1_array.go` | Arrays |
| `2_slice.go` | Slices |
| `3_map.go` | Maps |
| `4_pointer.go` | Pointers and pointer receivers |
| `5_struct.go` | Structs |
| `6_interface.go` | Interfaces |
| `7_chanel.go` | Channels |
| `7_gorutine.go` | Goroutines |
| `8_defer.go` | Deferred function calls |
| `9_read.go` | Reading a CSV file |
| `10_write.go` | Writing files |
| `11_handle.go` | HTTP handling |
| `12_jsonMarshal.go` | Go values to JSON with `json.Marshal` |
| `13_jsonUnMarshal.go` | JSON to Go values with `json.Unmarshal` |

## CSV example

`products.csv` contains 10 sample products with prices. The reader expects to
be run from the same directory as the CSV file:

```sh
cd go_basic_borntodev
go run 9_read.go
```

The writer creates `data.txt` and `name` in the current working directory:

```sh
go run 10_write.go
```

## Notes

- Do not run all example files together because several files define their own
  `main` function.
- The filenames `7_chanel.go`, `7_gorutine.go`, and `13_jsonUnMarshal.go`
  follow the current project names, although the conventional spellings are
  `channel`, `goroutine`, and `jsonUnmarshal`.
- The examples are for learning and are not production-ready. A production
  HTTP service should add HTTPS, request validation, timeouts, structured
  logging, monitoring, and rate limiting.

## Learning resources

- [Kongruksiam Go tutorial](https://youtu.be/pytqhPDTjnQ?si=ny6qhAodjOly_mJG)
- [BornToDev Go tutorial](https://youtu.be/fjEB75Xotxc?si=pJxwbRGYnWqfOg5n)
