// Package service provides a wrapper around NATS for microservices.
package service

import (
	"context"
	"errors"
	"log"
	"os"
	"time"

	"github.com/nats-io/nats.go"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/nats-io/nats.go/micro"
)

// NatsMicro is a wrapper around nats.Conn that provides a microservice interface.

type TinyService struct {
	nc     *nats.Conn
	svc    micro.Service
	sigs   chan os.Signal
	done   chan bool
	config jetstream.KeyValue

	groupName string
	group     micro.Group

	// Config is the configuration for the NatsMicro instance.
	micro.Config
}

// TinyServiceOption is a function that configures the NatsMicro instance.
type TinyServiceOption func(*TinyService)

// WithNatsConn sets the NATS connection for the NatsMicro instance.
func WithNatsConn(nc *nats.Conn) TinyServiceOption {
	return func(nm *TinyService) {
		nm.nc = nc
	}
}

// WithNatsURL sets the NATS URL for the NatsMicro instance.
func WithNatsURL(url string) TinyServiceOption {
	return func(nm *TinyService) {
		nc, err := nats.Connect(url)
		if err != nil {
			panic(err)
		}
		nm.nc = nc
	}
}

// WithNatsOptions sets the NATS options for the NatsMicro instance.
func WithNatsOptions(opts ...nats.Option) TinyServiceOption {
	return func(nm *TinyService) {
		nc, err := nats.Connect(nm.nc.ConnectedUrl(), opts...)
		if err != nil {
			panic(err)
		}
		nm.nc = nc
	}
}

// WithName sets the name for the NatsMicro instance.
func WithName(name string) TinyServiceOption {
	return func(nm *TinyService) {
		nm.Name = name
	}
}

// WithVersion sets the version for the NatsMicro instance.
func WithVersion(version string) TinyServiceOption {
	return func(nm *TinyService) {
		nm.Version = version
	}
}

// WithDescription sets the description for the NatsMicro instance.
func WithDescription(description string) TinyServiceOption {
	return func(nm *TinyService) {
		nm.Description = description
	}
}

// WithEndpoint sets the endpoint for the NatsMicro instance.
func WithEndpoint(endpoint *micro.EndpointConfig) TinyServiceOption {
	return func(nm *TinyService) {
		nm.Endpoint = endpoint
	}
}

// WithErrorHandler sets the error handler for the NatsMicro instance.
func WithErrorHandler(handler micro.ErrHandler) TinyServiceOption {
	return func(nm *TinyService) {
		nm.ErrorHandler = handler
	}
}

// WithDoneHandler sets the done handler for the NatsMicro instance.
func WithDoneHandler(handler micro.DoneHandler) TinyServiceOption {
	return func(nm *TinyService) {
		nm.DoneHandler = handler
	}
}

// WithStatsHandler sets the stats handler for the NatsMicro instance.
func WithStatsHandler(handler micro.StatsHandler) TinyServiceOption {
	return func(nm *TinyService) {
		nm.StatsHandler = handler
	}
}

// WithGroup sets the group name for the NatsMicro instance.
func WithGroup(groupName string) TinyServiceOption {
	return func(nm *TinyService) {
		nm.groupName = groupName
	}
}

// NewTinyService creates a new NatsMicro instance with the given NATS connection and configuration.
func NewTinyService(opts ...TinyServiceOption) (*TinyService, error) {
	nm := &TinyService{}
	nm.sigs = make(chan os.Signal, 1)
	nm.done = make(chan bool, 1)

	// Default NATS connection

	for _, opt := range opts {
		opt(nm)
	}

	if nm.nc == nil {
		return nil, nats.ErrNoServers
	}

	if nm.Name == "" {
		nm.Name = "NatsMicro"
	}

	if nm.Version == "" {
		nm.Version = "0.0.1"
	}
	if nm.Description == "" {
		nm.Description = "Nats Micro Service"
	}

	return nm, nil
}

func (nm *TinyService) AddEndpoint(name string, handler micro.Handler, opts ...micro.EndpointOpt) error {
	if nm.svc == nil {
		return nats.ErrNoServers
	}
	if nm.group != nil {
		return nm.group.AddEndpoint(name, handler, opts...)
	}
	return nm.svc.AddEndpoint(name, handler, opts...)
}

func (nm *TinyService) Init() error {
	if nm.svc != nil {
		return errors.New("service already initialized")
	}
	var err error
	nm.svc, err = micro.AddService(nm.nc, nm.Config)
	if err != nil {
		log.Printf("Error initializing service: %v", err)
		return err
	}
	log.Println("Service ID: ", nm.svc.Info().ID)
	if nm.groupName != "" {
		nm.group = nm.svc.AddGroup(nm.groupName)
		log.Printf("Group %s created", nm.groupName)
	}
	return nm.CreateConfig()
}

func (nm *TinyService) RunBlocking() error {
	if nm.svc == nil {
		return errors.New("service not initialized")
	}
	if nm.nc == nil {
		return errors.New("nats connection not initialized")
	}
	sig := <-nm.sigs
	log.Printf("Received signal: %s. Shutting down service...", sig)
	nm.Stop()
	log.Println("Service stopped")
	nm.done <- true
	return nil
}

func (nm *TinyService) Stop() {
	if nm.svc != nil {
		err := nm.svc.Stop()
		if err != nil {
			log.Printf("Error stopping service: %v", err)
		}
	}
}

func (nm *TinyService) CreateConfig() error {
	if nm.svc == nil {
		return nil
	}
	if nm.nc == nil {
		return nil
	}
	js, err := jetstream.New(nm.nc)
	if err != nil {
		log.Println("Error creating JetStream: ", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	kv, err := js.CreateOrUpdateKeyValue(ctx, jetstream.KeyValueConfig{
		Bucket: nm.configBucket(),
	})
	if err != nil {
		log.Println("Error creating KeyValue store: ", err)
	}
	log.Println("Config store created: ", kv.Bucket())
	nm.config = kv
	return nil
}

func (nm *TinyService) configBucket() string {
	return "tiny_config_" + nm.Name
}

func (nm *TinyService) ConfigStore() (jetstream.KeyValue, error) {
	if nm.config == nil {
		return nil, errors.New("config not initialized")
	}
	return nm.config, nil
}
