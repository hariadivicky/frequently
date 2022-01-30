# Frequently, the words counter

## Installation
requires go 1.13 or later

1. clone this repository
2. download dependencies: `go mod download`
3. build using `make build` or `go build -o frequently main.go`

## Usage

Main commands

| command | description                      |
|---------|----------------------------------|
| serve   | start http server                |
| req     | make request as client to server |
| help    | help about any command           |

### Serve Command

use `serve` command to start http server.

```sh
Usage:
  ./frequently serve [flags]

Flags:
  -a, --address string   http listening port (default ":8000")
  -h, --help             help for serve
```

### Req Command

use `req` command to make request as client to server

```sh
Usage:
  ./frequently req [flags]

Flags:
  -f, --file string   file to check
  -h, --help          help for req
  -i, --insensitive   insensitive mode
  -m, --max string    max result (default "10")
  -u, --url string    url for request (default "http://localhost:8000")
```

example to get top 5 words appearance from server: `./frequently -f ./data/seed_test.txt --max=5`