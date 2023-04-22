package firebase

import (
	"context"

	"github.com/sethvargo/go-envconfig"
)

// Config represents the configuration settings for the Firebase app.
type Config struct {
	Credentials string `env:"GOOGLE_APPLICATION_CREDENTIALS"`
}

// NewConfig reads Firebase configuration settings from environment variables.
func NewConfig() (cfg Config, err error) {
	err = envconfig.Process(context.Background(), &cfg)
	return
}
