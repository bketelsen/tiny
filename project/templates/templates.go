// Package templates provides templates for generating code
package templates

func TypeTemplate() []byte {
	return []byte(`package {{.Module}}

{{range $m,$v := .Service.MessageMap }}// {{ $m }} is a struct for the {{ $m }} type
type {{ $m }} struct { {{range $f, $fv := $v.FieldMap}}
  {{$fv.DeclarationName}} {{$fv.DeclarationType}} {{$fv.DeclarationTag}}{{end}}
}
{{end}}
{{range $m,$v := .Service.EnumMap}}// {{ $m }} is a type for the {{ $m }} enum
type {{$v.Name}} int

const ({{range $v.Values}}
	{{.Key}} {{$v.Name}} = {{.Value }}{{end}}
)
{{end}}
`)
}

func HandlerTemplate() []byte {
	return []byte(`// Package handlers contains the implementation of the {{.Service.Name}} service
package handlers
	
import (
	"{{.Module}}"

  "encoding/json"
  "log"

 	"github.com/nats-io/nats.go"
  "github.com/nats-io/nats.go/micro"



)

// {{.Endpoint.Name}} is a struct for the {{.Endpoint.Name}} endpoint
// It is the server implementation of the {{.Endpoint.Name}}Server interface
// TODO: Add fields to the struct if needed for server dependencies and state
type {{.Endpoint.Name}} struct {
  nc *nats.Conn
}

{{ $server := .Endpoint.Name }}{{$module := .Module}}{{range .Endpoint.GetAllMethods}}
 // {{.Name}} is the implementation of the {{$server}}.{{.Name}} endpoint
func (s *{{$server}}) {{.Name}}( req micro.Request )  {
  // Unmarshal the request
  input := &{{$module}}.{{.RequestTypeName}}{}
  err := json.Unmarshal(req.Data(),input)
  if err != nil {
    log.Println("Error unmarshalling request: ", err)
    return
  }
  
  // Create the response
  rsp := &{{$module}}.{{.ResponseTypeName}}{}
  // TODO: implement the endpoint logic
	err = req.RespondJSON(rsp)
	if err != nil {
		log.Println("Error responding:", err)
		return
	}
	return 
}{{end}}

// New{{.Endpoint.Name}} creates a new {{.Endpoint.Name}} struct
// TODO: Add parameters to the the function if needed to set server dependencies and state
func New{{.Endpoint.Name}}(nc *nats.Conn) *{{.Endpoint.Name}} {
	return &{{.Endpoint.Name}}{
    nc: nc,
  }
}
	`)
}

func ServiceTemplate() []byte {
	return []byte(`package main

import (
	"{{.Module}}/handlers"
	"log"
	"os"
	"strings"

  "github.com/bketelsen/tiny/service"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"

)

func main() {

	url, exists := os.LookupEnv("NATS_URL")
	if !exists {
		url = nats.DefaultURL
	} else {
		url = strings.TrimSpace(url)
	}

	if strings.TrimSpace(url) == "" {
		url = nats.DefaultURL
	}

	// Connect to the server
	nc, err := nats.Connect(url)
	if err != nil {
		log.Fatal(err)
		return
	}
	nm, err := service.NewTinyService(
		service.WithNatsConn(nc),
		service.WithName("{{.Service.Name}}"),
		service.WithVersion("0.0.1"),
		service.WithDescription("{{.Service.Description}}"),
    service.WithGroup("{{.Service.Name}}"),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = nm.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Service initialized")

  {{range .Service.GetAllEndpoints}}
  // {{.Name}} handler
  {{.ClientStructName  }}Handler := handlers.New{{.Name}}(nc)
	// register {{.Name}}Handler
  {{$service := . }}{{range .GetAllMethods}}
	nm.AddEndpoint("{{$service.Name}}{{.Name}}", micro.HandlerFunc({{$service.ClientStructName}}Handler.{{.Name}})){{end}}{{end}}

  err = nm.RunBlocking()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Service stopped")
}
`)
}

func ServiceHandlerTemplate() []byte {
	return []byte(`package handlers

import (
	"context"

	"{{.Module}}"
)

{{ $server := .Endpoint.Name }}func (s *{{$server}}) {{.Def.Name}}(ctx context.Context, req *{{.Module}}.{{.Def.Request.String}}, rsp *{{.Module}}.{{.Def.Response.String}}) error {

	return nil
}
`)
}

func ServiceClientTemplate() []byte {
	return []byte(`// Package {{.Module}} defines the types and interfaces for the {{.Service.Name}} service
package {{.Module}}

import (
	"encoding/json"

	"github.com/nats-io/nats.go"
)


{{$sn := .Service.Name}}{{range .Service.GetAllEndpoints}}
// {{.Name}} Methods
{{$endpoint := .Name}}
{{range .GetAllMethods}}
func {{$endpoint}}{{.Name}}(nc *nats.Conn, in {{.RequestTypeName}}) ({{.ResponseTypeName}}, error){
	bb, err := json.Marshal(in)
	if err != nil {
		return {{.ResponseTypeName}}{}, err
	}
  var out {{.ResponseTypeName}}
  resp, err := nc.Request("{{$sn}}.{{$endpoint}}{{.Name}}", bb, nats.DefaultTimeout)
  if err != nil {
    return {{.ResponseTypeName}}{}, err
  }
    	err = json.Unmarshal(resp.Data, &out)
	if err != nil {
		return {{.ResponseTypeName}}{}, err
	}
	return out, nil
}
{{end}}
{{end}}

`)
}

func ConfigTemplate() []byte {
	return []byte(`service="{{.Service}}"
description="replace with a description of the service"

type {{.Method}}Request {
  input string
}

type {{.Method}}Response {
  output string
}

endpoint {{.Endpoint}} {
  rpc {{.Method}}({{.Method}}Request) returns ({{.Method}}Response)
}
`)
}

func GitIgnoreTemplate() []byte {
	return []byte(`.DS_Store  
# If you prefer the allow list template instead of the deny list, see community template:
# https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
#
# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib

# Test binary, built with go test -c
*.test

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# Dependency directories (remove the comment below to include it)
# vendor/

# Go workspace file
go.work
go.work.sum

# env file
.env

# The binary output of the build tool
/{{.SERVICE_NAME}}

# Taskfile
/.task

  `)
}

func TaskfileTemplate() []byte {
	return []byte(`# https://taskfile.dev

version: "3"

env:
  GO111MODULE: on
  GOPROXY: https://proxy.golang.org,direct

tasks:

  setup:
    desc: Install dependencies
    cmds:
      - go mod tidy

  build:
    desc: Build the binary
    sources:
      - ./**/*.go
    generates:
      - ./{{.SERVICE_NAME}}
    cmds:
      - go build ./cmd/{{.SERVICE_NAME}}

  install:
    desc: Install the binary locally
    sources:
      - ./**/*.go
    cmds:
      - go install ./cmd/{{.SERVICE_NAME}} 

  test:
    desc: Run tests
    cmds:
      - go test -failfast -race -coverpkg=./... -covermode=atomic -coverprofile=coverage.txt ./...  -timeout=15m

  cover:
    desc: Open the cover tool
    cmds:
      - go tool cover -html=coverage.txt

  ci:
    desc: Run all CI steps
    cmds:
      - task: build
      - task: test

  default:
    desc: Runs the default tasks
    cmds:
      - task: ci

  run:
    desc: Run the service
    deps:
      - build
    cmds:
      - ./{{.SERVICE_NAME}}

  clean:
    desc: Clean the project	
    cmds:
      - rm ./{{.SERVICE_NAME}}

`)
}
