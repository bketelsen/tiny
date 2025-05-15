# tiny

[![Go Reference](https://pkg.go.dev/badge/github.com/bketelsen/tiny.svg)](https://pkg.go.dev/github.com/bketelsen/tiny)

A developer toolkit for building NATS-based microservices in Go.

## Overview

tiny provides tools to streamline the development of microservices that use NATS for communication. It offers code generation capabilities from service definitions written in a MuCL (Micro Control Language) format.

## Installation

```bash
go install github.com/bketelsen/tiny@latest
```

## Quick Start

1. Initialize a new project:

```bash
tiny init
```

2. Generate the service code:

```bash
tiny gen
```

3. Implement your business logic in the generated files

## Commands

### Generate

Generate code from a MuCL file:

```bash
tiny gen --definition <file.tiny> [--types] [--force]
```

Options:

- `--types`: Generate only type definitions
- `--force`: Overwrite existing files

### Initialize

Create a new project structure:

```bash
tiny init [project-name]
```

## Documentation

Comprehensive documentation is available at the [tiny documentation site](https://bketelsen.github.io/tiny/).

- [Installation Guide](https://bketelsen.github.io/tiny/guide/installation)
- [Getting Started](https://bketelsen.github.io/tiny/guide/getting-started)
- [MuCL Reference](https://bketelsen.github.io/tiny/reference/mucl)

## Development

### Prerequisites

- Go 1.24 or newer
- [Task](https://taskfile.dev) for development workflows

### Build from source

```bash
task build
```

### Run tests

```bash
task test
```

### Generate documentation

```bash
task docs:build
```

Preview documentation locally:

```bash
task docs:preview
```

## Contributing

Contributions are welcome! Please see our [Contributing Guide](https://github.com/bketelsen/tiny/blob/main/CONTRIBUTING.md) for more details.

## License

This project is licensed under the MIT License.
