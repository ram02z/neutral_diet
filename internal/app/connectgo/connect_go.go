// Package connectgo implements utilities for building and serving neutral-diet services.
package connectgo

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"firebase.google.com/go/messaging"
	"github.com/bufbuild/connect-go"
	grpchealth "github.com/bufbuild/connect-grpchealth-go"
	"github.com/justinas/alice"
	"github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1/foodv1connect"
	"github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/job/v1/jobv1connect"
	"github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1/userv1connect"
	"github.com/ram02z/neutral_diet/internal/service"
	"github.com/ram02z/neutral_diet/internal/service/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// RegisterConnectGoServerInput represents the input to the RegisterConnectGoServer function.
type RegisterConnectGoServerInput struct {
	Logger     *zerolog.Logger
	ConnectSvc *service.ConnectWrapper
	Mux        *http.ServeMux `name:"connectGoMux"`
	AuthClient *auth.Client
}

// RegisterConnectGoServer registers the handlers for the connect-go server.
func RegisterConnectGoServer(in RegisterConnectGoServerInput) {
	api := http.NewServeMux()
	api.Handle(foodv1connect.NewFoodServiceHandler(
		in.ConnectSvc,
		connect.WithInterceptors(
			connectInterceptorForLogger(in.Logger),
		),
	))
	api.Handle(userv1connect.NewUserServiceHandler(
		in.ConnectSvc,
		connect.WithInterceptors(
			connectInterceptorForLogger(in.Logger),
			connectInterceptorForUserAuth(in.AuthClient),
		),
	))
	api.Handle(jobv1connect.NewJobServiceHandler(
		in.ConnectSvc,
		connect.WithInterceptors(
			connectInterceptorForLogger(in.Logger),
			connectInterceptorForCloudSchedulerAuth(),
		),
	))
	checker := grpchealth.NewStaticChecker(
		"neutral_diet.food.v1.FoodService",
		"neutral_diet.user.v1.UserService",
		"neutral_diet.job.v1.JobService",
	)
	in.Mux.Handle(grpchealth.NewHandler(checker))
	in.Mux.Handle("/api/", http.StripPrefix("/api", api))
}

// A Server represents a HTTP server with a shutdown timer.
type Server struct {
	Server          *http.Server
	notify          chan error
	Mux             *http.ServeMux
	ShutdownTimeout time.Duration
}

// NewConnectWrapper creates a [service.ConnectWrapper] instance
func NewConnectWrapper(
	s *db.Store,
	a *auth.Client,
	m *messaging.Client,
) *service.ConnectWrapper {
	return service.NewConnectWrapper(s, a, m)
}

// NewConnectGoServer starts and returns server.
// Adds HTTP logger handlers to Mux.
func NewConnectGoServer(
	logger *zerolog.Logger,
	cfg Config,
) *Server {
	mux := http.NewServeMux()
	address := fmt.Sprintf(":%d", cfg.ConnectConfig.Port)

	c := alice.New()
	c = c.Append(hlog.NewHandler(*logger))
	// Thanks to that handler, all our logs will come with some prepopulated fields.
	c = c.Append(hlog.AccessHandler(func(
		r *http.Request,
		status, size int,
		duration time.Duration,
	) {
		hlog.FromRequest(r).Info().
			Str("method", r.Method).
			Stringer("url", r.URL).
			Int("status", status).
			Int("size", size).
			Dur("duration", duration).
			Msg("")
	}))
	c = c.Append(hlog.RemoteAddrHandler("ip"))
	c = c.Append(hlog.UserAgentHandler("user_agent"))
	c = c.Append(hlog.RefererHandler("referer"))
	c = c.Append(hlog.RequestIDHandler("req_id", "Request-Id"))
	c = c.Append(newCORS().Handler)

	// Here is your final handler
	h := c.Then(mux)

	srv := &http.Server{
		Addr: address,
		// Use h2c so we can serve HTTP/2 without TLS.
		Handler: h2c.NewHandler(
			h,
			&http2.Server{},
		),
	}

	s := &Server{
		Server:          srv,
		Mux:             mux,
		notify:          make(chan error, 1),
		ShutdownTimeout: cfg.ConnectConfig.ShutdownTimeout,
	}

	s.start(logger)

	return s
}

// start listen and serves server in separate goroutine
func (s *Server) start(logger *zerolog.Logger) {
	go func() {
		logger.Info().Str("address", s.Server.Addr).Str("path", "/api").Msg("Listening for connect-go")
		s.notify <- s.Server.ListenAndServe()
		close(s.notify)
	}()
}

// Notify returns the notify channel
func (s *Server) Notify() <-chan error {
	return s.notify
}

// Shutdown shuts down the server after all connections active are finished
func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()

	return s.Server.Shutdown(ctx)
}
