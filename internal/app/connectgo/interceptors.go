package connectgo

import (
	"context"

	"github.com/bufbuild/connect-go"
	"github.com/rs/zerolog"
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
