package firebase

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

type Config struct {
	Credentials string `env:"FIREBASE_CONFIG,required"`
}

func VerifyConfig() (err error) {
	var cfg Config
	err = envconfig.Process(context.Background(), &cfg)
	return err
}
