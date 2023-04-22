// Package service implements the RPCs for neutral-diet services
package service

import (
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/service/db"
)

// A ConnectWrapper represents the PostgreSQL database, Firebase Authentication and Firebase
// messaging clients.
type ConnectWrapper struct {
	s *db.Store
	a *auth.Client
	m *messaging.Client
}

// RequestWithValidator is an abstract interface used for connect-go requests.
type RequestWithValidator interface {
	Validate() error
}

// NewConnectWrapper creates a new instance of [ConnectWrapper].
func NewConnectWrapper(s *db.Store, a *auth.Client, m *messaging.Client) *ConnectWrapper {
	return &ConnectWrapper{s: s, a: a, m: m}
}

func validate(r RequestWithValidator) error {
	err := r.Validate()
	if err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}

	return nil
}
