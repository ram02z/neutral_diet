package auth

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

func NewAuth(cfg Config) (*auth.Client, error) {
	opt := option.WithCredentialsFile(cfg.Credentials)
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}

	auth, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing auth client: %v", err)
	}

	return auth, nil
}
