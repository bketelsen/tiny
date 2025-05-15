# Getting Started

## Install the tiny CLI tool

Visit [Install](/guide/installation) for more details.

## Initialize a new tiny specification

Run `tiny init` and answer a few questions. It creates a new directory and generates a sample service definition file called `service.tiny`.

## Edit the service definition


## Generate the project

Edit the service definition to suit your needs, then save it and generate your project:

```sh
tiny gen
```

Tiny generates all the boilerplate for your project:

```
❯ tree
.
├── client.go
├── cmd
│   └── search
│       └── main.go
├── config.go
├── Dockerfile
├── go.mod
├── go.sum
├── handlers
│   └── searchservice.go
├── search
├── service.tiny
├── Taskfile.yml
└── types.go

4 directories, 11 files
```