package main

import (
	"log"
	"os"
	"strings"

	"_example/handlers"

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
		service.WithName("users"),
		service.WithVersion("0.0.1"),
		service.WithDescription("User Management Service"),
		service.WithGroup("users"),
	)
	if err != nil {
		log.Fatal(err)
	}

	err = nm.Init()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Service initialized")

	// User handler
	userHandler := handlers.NewUser(nc)
	// register UserHandler

	nm.AddEndpoint("UserGet", micro.HandlerFunc(userHandler.Get))
	nm.AddEndpoint("UserUnlock", micro.HandlerFunc(userHandler.Unlock))

	err = nm.RunBlocking()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Service stopped")
}
