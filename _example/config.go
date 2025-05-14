package _example

type Config struct {
	// Cache is a struct for Cache configuration
	Cache struct {
		Host string `json:"host" yaml:"host" env:"CACHE_HOST" env-description:"Cache host"`
		Port int    `json:"port" yaml:"port" env:"CACHE_PORT" env-description:"Cache port"`
	} `json:"cache" yaml:"cache"`

	NatsURL string `env:"NATS_URL" env-description:"NATS URL" json:"nats_url" yaml:"nats_url"`
}
