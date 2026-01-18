# osxec

[![Go Reference](https://pkg.go.dev/badge/github.com/sitnikovik/osxec.svg)](https://pkg.go.dev/github.com/sitnikovik/osxec)
[![Go Report Card](https://goreportcard.com/badge/github.com/sitnikovik/osxec)](https://goreportcard.com/report/github.com/sitnikovik/osxec)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/sitnikovik/osxec)
[![Release](https://img.shields.io/github/v/release/sitnikovik/osxec?style=flat)](https://github.com/sitnikovik/osxec/releases)

Golang toolkit for robust, object-oriented execution and management of OS shell commands with context support, output handling, and error management.

## Features

- Simple API for executing shell commands in **OOP style**
- Pass command arguments and environment variables
- **Synchronous** and **asynchronous** execution with context.Context support
- Capture and process stdout/stderr output
- Retrieve and handle process exit codes
- Robust error handling for command execution
- Minimal dependencies, cross-platform (works on all Go-supported OSes)
- Clean, extensible interfaces for integration into other projects

## Installation

```bash
go get github.com/sitnikovik/osxec
```

## Getting started

```go
import (
    "context"
    "log"

    "github.com/sitnikovik/osxec/command"
    "github.com/sitnikovik/osxec/process"
    "github.com/sitnikovik/osxec/shell"
)

func main() {
    ctx := context.Background()
    res := process.
        NewProcess(
            shell.NewShell(),
            command.NewCommand("echo", "Hello, World!"),
        ).
        Execution(ctx)
    if err := res.Err(); err != nil {
        log.Fatalf("Process execution failed: %v", err)
        return
    }
    log.Printf("Process output as string: '%s'\n", res.String())
}
```

## Documentation

Full API reference is available on [pkg.go.dev](https://pkg.go.dev/github.com/sitnikovik/osxec)

## Requirements

- Go 1.23.4 (as declared in `go.mod`) — newer Go versions should work
- No external runtime dependencies

## License

This project is licensed under the MIT License — see the [LICENSE](LICENSE) file for details.

## Author / Contact

Maintained by [Ilya Sitnikov](https://github.com/sitnikovik)
