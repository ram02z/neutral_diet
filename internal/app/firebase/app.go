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

func NewApp(cfg Config) (*FirebaseApp, error) {
	if cfg.Credentials != "" {
		opt := option.WithCredentialsFile(cfg.Credentials)
		app, err := firebase.NewApp(context.Background(), nil, opt)
		if err != nil {
			return nil, fmt.Errorf("error initializing app with credentials file: %v", err)
		}
		return &FirebaseApp{
			app: app,
		}, nil
	} else {
		app, err := firebase.NewApp(context.Background(), nil)
		if err != nil {
			return nil, fmt.Errorf("error initializing app: %v", err)
		}
		return &FirebaseApp{
			app: app,
		}, nil
	}
}
