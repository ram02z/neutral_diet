package connectgo

import (
	"context"
	"errors"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog"
	"github.com/ram02z/neutral_diet/internal"
)

func getUnaryInterceptors(logger *zerolog.Logger) []connect.Interceptor {
	return []connect.Interceptor{
		connectInterceptorForLogger(logger),
	}
}

func connectInterceptorForLogger(logger *zerolog.Logger) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			return next(logger.WithContext(ctx), req)
		})
	}
	return connect.UnaryInterceptorFunc(interceptor)
}

func connectInterceptorForAuth(auth *auth.Client) connect.UnaryInterceptorFunc {
	interceptor := func(next connect.UnaryFunc) connect.UnaryFunc {
		return connect.UnaryFunc(func(
			ctx context.Context,
			req connect.AnyRequest,
		) (connect.AnyResponse, error) {
			// Get raw header
			authHeader := req.Header().Get("Authorization")
			if authHeader == "" {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("no authorization token provided"),
				)
			}

			// Confirm request is sending Bearer authentication credentials
			if !strings.HasPrefix(authHeader, "Bearer ") {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("authorization token is malformed"),
				)
			}

			// Verify token
			token, err := auth.VerifyIDToken(ctx, authHeader[7:])
			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					errors.New("authorization token could not be verified"),
				)
			}
			req.Header().Set(config.UserIDHeaderKey, token.UID)

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}

