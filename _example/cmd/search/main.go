package main

import (
	"_example/handlers"
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
		service.WithName("search"),
		service.WithVersion("0.0.1"),
		service.WithDescription("something with spaces"),
		service.WithGroup("search"),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = nm.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Service initialized")

	// SearchService handler
	searchServiceHandler := handlers.NewSearchService(nc)
	// register SearchServiceHandler

	nm.AddEndpoint("Search", micro.HandlerFunc(searchServiceHandler.Search))

	err = nm.RunBlocking()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Service stopped")
}
