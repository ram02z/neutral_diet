package firebase

import (
	"context"
	"fmt"

	"firebase.google.com/go/messaging"
)

// NewMessaging returns instance of Firebase Messaging client.
func (f *FirebaseApp) NewMessaging() (*messaging.Client, error) {
	messaging, err := f.app.Messaging(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error initializing messaging client: %v", err)
	}

	return messaging, nil
}
