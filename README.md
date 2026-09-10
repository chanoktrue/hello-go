# Hello Go

A small Go HTTP API that responds with a greeting.

## Requirements

- Go 1.27.1 or newer

## Run locally

```bash
go run .
```

The server listens on `http://localhost:8080`.

## API

### `GET /hello`

Returns:

```text
Hello from Go on Omarchy!
```

Try it with:

```bash
curl http://localhost:8080/hello
```

Other HTTP methods return `405 Method Not Allowed`.

## Build

```bash
go build -o hello-go .
./hello-go
```

## Test

Run the project tests with:

```bash
go test ./...
```

## Project structure

```text
.
├── go.mod
├── main.go
└── README.md
```
