# Hello Go

A practical learning space for Go, with examples that can be read, run, and
extended. Topics range from language fundamentals, JSON, and file I/O to
concurrency, HTTP APIs, middleware, and CORS.

## Why Go?

Go is a strong choice for developers who want to build backend systems that
are fast, simple, and reliable. Key reasons to start learning Go include:

- **Concise, readable syntax** keeps the focus on solving problems rather than language complexity.
- **Simple compilation and deployment** allow applications to be distributed as a single binary.
- **Strong performance** makes Go suitable for web services, APIs, CLIs, cloud services, and high-throughput systems.
- **Built-in concurrency** with goroutines and channels provides a clear way to learn concurrent programming.
- **A powerful standard library** includes practical support for HTTP, JSON, file I/O, testing, and system programming.
- **Team-friendly conventions** make Go code straightforward to review, maintain, and extend.

## Motivation

This repository was inspired by the work and guidance of
[Jaime Still](https://github.com/JaimeStill), a software engineer who regularly
explores and builds with Go. This approach shows that effective learning does
not require starting with a large system. It is better to begin with small,
understandable programs and gradually grow them into packages, services, and
usable APIs.

The goal of this repository is hands-on learning:

1. Understand the language fundamentals and standard library.
2. Practice separating responsibilities into modules and packages.
3. Build HTTP APIs and experiment with requests and responses using `curl`.
4. Learn to read errors, test code, and improve structure incrementally.

This repository is a learning project, not a production-ready service. The
examples are structured as a foundation that can be extended toward real-world
systems.

## Requirements

- Go 1.26.4 for the main learning modules
- Go 1.27.1 for the `tutorial` module
- macOS, Linux, or another platform supported by Go

Check the installed Go version:

```sh
go version
```

## Repository layout

| Module / directory | Purpose | Run from |
|---|---|---|
| `tutorial/` | Basic HTTP server tutorial | `tutorial/` |
| `go_basic_borntodev/` | Go syntax, JSON, file I/O, HTTP, middleware, and CORS examples | `go_basic_borntodev/` |
| `go_basic_kongruksiam/` | Go language fundamentals and package examples | `go_basic_kongruksiam/` |
| `go_api_chatgpt/` | In-memory product CRUD API using `net/http` | `go_api_chatgpt/` |
| `go_api_chatgpt_package/` | Product API organized into `handler` and `model` packages | `go_api_chatgpt_package/` |
| `go_CURD/` | Minimal `/hello` HTTP example | `go_CURD/` |
| `go_test_api/` | In-memory product CRUD API with validation and JSON handling | `go_test_api/` |

The repository root is not a Go module. Run Go commands from the module
directory being studied.

## Quick start

Run the standalone product CRUD API from its main example:

```sh
cd go_api_chatgpt
go run main.go
```

The API listens on `http://127.0.0.1:8080`.

In another terminal:

```sh
curl http://127.0.0.1:8080/products
```

Stop the server with `Ctrl+C`.

## Product API

The `go_api_chatgpt` service stores products in memory. Data is reset when the
process stops.

### Endpoints

| Method | Endpoint | Description |
|---|---|---|
| `GET` | `/products` | List all products |
| `GET` | `/products/{code}` | Get one product by code |
| `POST` | `/products` | Create a product |
| `PUT` | `/products/{code}` | Replace a product |
| `DELETE` | `/products/{code}` | Delete a product |

### Request examples

Get one product:

```sh
curl http://127.0.0.1:8080/products/P001
```

Create a product:

```sh
curl -X POST http://127.0.0.1:8080/products \
  -H "Content-Type: application/json" \
  -d '{"productCode":"P003","productName":"Keyboard","price":900}'
```

Update a product:

```sh
curl -X PUT http://127.0.0.1:8080/products/P001 \
  -H "Content-Type: application/json" \
  -d '{"productCode":"P001","productName":"Studio Speaker","price":2800}'
```

Delete a product:

```sh
curl -i -X DELETE http://127.0.0.1:8080/products/P001
```

## Learning examples

### `go_basic_borntodev`

Each example is a separate `package main` program. Run one file at a time:

```sh
cd go_basic_borntodev
go run 1_array.go
go run 12_jsonMarshal.go
go run 13_jsonUnMarshal.go
go run 14_workwithrequest.go
go run 15_middleware.go
go run 16_CORS.go
```

Important examples:

| File | Topic |
|---|---|
| `1_array.go` – `8_defer.go` | Core language, collections, functions, and concurrency |
| `9_read.go` | Read CSV data |
| `10_write.go` | Write local files |
| `11_handle.go` | Basic HTTP handling |
| `12_jsonMarshal.go` | Encode Go values as JSON |
| `13_jsonUnMarshal.go` | Decode JSON into Go values |
| `14_workwithrequest.go` | Product CRUD API with validation |
| `15_middleware.go` | HTTP middleware |
| `16_CORS.go` | CORS and preflight requests |

The product API in `14_workwithrequest.go` uses port `8000` and the `/product`
path:

```sh
cd go_basic_borntodev
go run 14_workwithrequest.go
curl http://127.0.0.1:8000/product
```

### Other HTTP examples

Run the minimal server in `go_CURD`:

```sh
cd go_CURD
go run .
curl http://127.0.0.1:8000/hello
```

Run the validated product API in `go_test_api`:

```sh
cd go_test_api
go run .
curl http://127.0.0.1:8080/products
```

## Testing and verification

Run tests for a specific module:

```sh
cd go_api_chatgpt
go test main.go
go vet main.go
```

Some learning examples in the same module define their own `main` function.
Run or test those files individually. For package-based modules, use
`go test ./...` from that module directory. The repository root does not
support `go test ./...` because it has no root `go.mod`.

If the default Go build cache is unavailable on a managed or synced
filesystem, use a temporary cache:

```sh
GOCACHE=/private/tmp/hello-go-gocache go test ./...
```

## Development notes

- Keep examples isolated; several files define their own `main` function.
- Product APIs use in-memory data and are for learning, not production use.
- Production services should add persistent storage, request validation,
  graceful shutdown, timeouts, structured logging, monitoring, and rate
  limiting.
- Do not commit credentials, tokens, private keys, or local machine data.
- `go_api_chatgpt_package` is an experimental package-based refactor with
  separate `handler` and `model` packages. Keep `model` independent from
  `handler` and avoid importing a package from itself.

## Inspiration and learning resources

The learning path in this repository is inspired by Jaime Still and supported
by the following resources:

- [Jaime Still on GitHub](https://github.com/JaimeStill)
- [Kongruksiam Go tutorial](https://youtu.be/pytqhPDTjnQ?si=ny6qhPDTjnWmJG)
- [BornToDev Go tutorial](https://youtu.be/fjEB75Xotxc?si=pJxwbRGYnWqfOg5n)
- [mikelopster Go tutorial](https://www.youtube.com/watch?v=KnwwdVBdmzg&list=PLwZ0y9k-cYXAJESl_kMGMQtXaSYpRV5U2)
- [Go Web Service reference](https://github.com/standards-lab/go-web-service)

Advice for beginners: do not try to learn everything at once. Start with the
basic syntax, write small programs, understand the errors you encounter, and
then progress to packages, HTTP APIs, testing, and system design. Consistent
practice and regular code review are what turn knowledge into skill.

## License

This repository is a personal learning project. Add a project-specific
license before distributing it as a reusable library or production service.
