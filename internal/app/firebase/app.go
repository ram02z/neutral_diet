package firebase

import (
	"context"
	"fmt"

	firebase "firebase.google.com/go"
)

type FirebaseApp struct {
	app *firebase.App
}

func NewApp() (*FirebaseApp, error) {
	app, err := firebase.NewApp(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	return &FirebaseApp{
		app: app,
	}, nil
}
