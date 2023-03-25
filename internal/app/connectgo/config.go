package connectgo

import (
	"context"
	"time"

	envconfig "github.com/sethvargo/go-envconfig"
)

type Config struct {
	ConnectConfig *ConnectConfig `env:",prefix=GRPC_CONNECT_"`
}

type ConnectConfig struct {
	Port            int           `env:"PORT,default=8080"`
	ShutdownTimeout time.Duration `env:"TIMEOUT,default=3s"`
}

func NewConfig() (cfg Config, err error) {
	err = envconfig.Process(context.Background(), &cfg)
	return
}
