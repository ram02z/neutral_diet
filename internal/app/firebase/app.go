// Package firebase provides wrapper functions for Firebase App initialisers.
package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
	"google.golang.org/api/option"
)

type FirebaseApp struct {
	app *firebase.App
}

// NewApp creates a Firebase app with given configuration.
// If the Config Credentials field is non-empty, the field is used as the credentials file path.
// Otherwise, Google application default credentials are used.
func NewApp(cfg Config) (*FirebaseApp, error) {
	var app *firebase.App
	var err error
	if cfg.Credentials != "" {
		opt := option.WithCredentialsFile(cfg.Credentials)
		app, err = firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return nil, fmt.Errorf("error initializing app with credentials file: %v", err)
		}
	} else {
		app, err = firebase.NewApp(context.Background(), nil)
		if err != nil {
			return nil, fmt.Errorf("error initializing app: %v", err)
		}
	}
	return &FirebaseApp{app: app}, nil
}
