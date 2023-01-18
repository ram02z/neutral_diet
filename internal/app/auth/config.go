package auth

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Credentials string `env:"GOOGLE_APPLICATION_CREDENTIALS,required"`
}

func NewConfig() (cfg Config, err error) {
	err = envconfig.Process(context.Background(), &cfg)
	return
}
