package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/auth"
)

func (f *FirebaseApp) NewAuth() (*auth.Client, error) {
	auth, err := f.app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing auth client: %v", err)
	}

	return auth, nil
}
