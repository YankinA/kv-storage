# kv-storage

This is a simple in-memory key-value database implemented in Go.

## Features

- TCP protocol support
- Key-value storage

## Usage

To run the server:
```sh
go run main.go
```

## Supports commands:
  - `SET key value`
  - `GET key`
  - `DEL key`


## Installation

To run this project, install Go version 1.21 or later.

### Building the Binary

#### For Linux:

Run the following command to build the binary for Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o key-value-store-linux
```

#### For macOS:
```bash
GOOS=darwin GOARCH=amd64 go build -o key-value-store-macos
```