package connectgo

import (
	"context"
	"time"

	envconfig "github.com/sethvargo/go-envconfig"
)

// Config represents the configuration settings for the connect-go server.
type Config struct {
	ConnectConfig *ConnectConfig `env:",prefix=GRPC_CONNECT_"`
}

// ConnectConfig represents the connection settings for the connect-go server.
type ConnectConfig struct {
	Port            int           `env:"PORT,default=8080"`
	ShutdownTimeout time.Duration `env:"TIMEOUT,default=3s"`
}


// NewConfig reads connect-go configuration settings from environment variables.
func NewConfig() (cfg Config, err error) {
	err = envconfig.Process(context.Background(), &cfg)
	return
}
