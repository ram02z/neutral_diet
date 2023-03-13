package service

import (
	"firebase.google.com/go/auth"
	"github.com/bufbuild/connect-go"
	"github.com/ram02z/neutral_diet/internal/service/db"
)

type ConnectWrapper struct {
	s *db.Store
	a *auth.Client
}

type Validator interface {
	Validate() error
}

func NewConnectWrapper(s *db.Store, a *auth.Client) *ConnectWrapper {
	return &ConnectWrapper{s: s, a: a}
}

func validate(r Validator) error {
	err := r.Validate()
	if err != nil {
		return connect.NewError(connect.CodeInvalidArgument, err)
	}

	return nil
}
