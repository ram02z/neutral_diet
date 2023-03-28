package service

import (
	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/service/db"
)

type ConnectWrapper struct {
	s *db.Store
	a *auth.Client
	m *messaging.Client
}

type Validator interface {
	Validate() error
}

func NewConnectWrapper(s *db.Store, a *auth.Client, m *messaging.Client) *ConnectWrapper {
	return &ConnectWrapper{s: s, a: a, m: m}
}

func validate(r Validator) error {
	err := r.Validate()
	if err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}

	return nil
}
