package service

import (
	"context"
	"errors"
	"fmt"
	"net/http"

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

func (c *ConnectWrapper) verify(
  ctx context.Context,
	header http.Header,
) (*auth.Token, error) {
	accessToken := header.Get("X-ID-Token")
	if accessToken == "" {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			errors.New("request is missing 'X-ID-Token' header"),
		)
	}

	token, err := c.a.VerifyIDToken(ctx, accessToken)
	if err != nil {
		return nil, connect.NewError(
			connect.CodeInvalidArgument,
			fmt.Errorf("'accessToken' header is invalid: %v", err),
		)
	}

	return token, nil
}
