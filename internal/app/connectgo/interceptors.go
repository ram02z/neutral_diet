package connectgo

import (
	"context"
	"errors"
	"strings"

	"firebase.google.com/go/auth"
	"github.com/bufbuild/connect-go"
	config "github.com/ram02z/neutral_diet/internal"
	"github.com/rs/zerolog"
	"google.golang.org/api/idtoken"
)

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

func connectInterceptorForUserAuth(auth *auth.Client) connect.UnaryInterceptorFunc {
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

			newCtx := context.WithValue(ctx, config.UserIDKey, token.UID)

			return next(newCtx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}

func connectInterceptorForCloudSchedulerAuth() connect.UnaryInterceptorFunc {
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

			token := authHeader[7:]
			_, err := idtoken.Validate(ctx, token, "")
			if err != nil {
				return nil, connect.NewError(
					connect.CodeUnauthenticated,
					err,
				)
			}

			// TODO: verify email from Claims

			return next(ctx, req)
		})
	}

	return connect.UnaryInterceptorFunc(interceptor)
}
