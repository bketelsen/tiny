package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"_example"
	"_example/handlers"

	"github.com/bketelsen/tiny/cleanenv"
	"github.com/bketelsen/tiny/service"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
)

// Args command-line parameters
type Args struct {
	ConfigPath string
}

func main() {
	var cfg _example.Config

	args := ProcessArgs(&cfg)

	// read configuration from the file and environment variables
	if err := cleanenv.ReadConfig(args.ConfigPath, &cfg); err != nil {
		log.Printf("Error reading configuration from file: %v", err)
		err = cleanenv.ReadEnv(&cfg)
		if err != nil {
			log.Printf("Error reading configuration from environment variables: %v", err)
			return
		}
	}

	if strings.TrimSpace(cfg.NatsURL) == "" {
		cfg.NatsURL = nats.DefaultURL
	}

	// Connect to the server
	nc, err := nats.Connect(cfg.NatsURL)
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
	userHandler := handlers.NewUser(nc, nm, &cfg)
	// register UserHandler

	nm.AddEndpoint("UserGet", micro.HandlerFunc(userHandler.Get))
	nm.AddEndpoint("UserUnlock", micro.HandlerFunc(userHandler.Unlock))

	err = nm.RunBlocking()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Service stopped")
}

// ProcessArgs processes and handles CLI arguments
func ProcessArgs(cfg interface{}) Args {
	var a Args

	f := flag.NewFlagSet("users", 1)
	f.StringVar(&a.ConfigPath, "c", "config.yml", "Path to configuration file")

	fu := f.Usage
	f.Usage = func() {
		fu()
		envHelp, _ := cleanenv.GetDescription(cfg, nil)
		fmt.Fprintln(f.Output())
		fmt.Fprintln(f.Output(), envHelp)
	}

	f.Parse(os.Args[1:])
	return a
}
