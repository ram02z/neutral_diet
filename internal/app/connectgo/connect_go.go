package connectgo

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"firebase.google.com/go/auth"
	"github.com/bufbuild/connect-go"
	"github.com/justinas/alice"
	"github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/food/v1/foodv1connect"
	"github.com/ram02z/neutral_diet/internal/gen/idl/neutral_diet/user/v1/userv1connect"
	"github.com/ram02z/neutral_diet/internal/service"
	"github.com/ram02z/neutral_diet/internal/service/db"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/hlog"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type RegisterConnectGoServerInput struct {
	Logger     *zerolog.Logger
	ConnectSvc *service.ConnectWrapper
	Mux        *http.ServeMux `name:"connectGoMux"`
	AuthClient *auth.Client
}

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
			connectInterceptorForAuth(in.AuthClient),
		),
	))
	// checker := grpchealth.NewStaticChecker(
	// 	// protoc-gen-connect-go generates package-level constants
	// 	// for these fully-qualified protobuf service names, so we'd be able
	// 	// to reference foov1beta1.FooService as opposed to foo.v1beta1.FooService.
	// 	"coop.drivers.foov1beta1.FooService",
	// )
	// in.Mux.Handle(grpchealth.NewHandler(checker))
	in.Mux.Handle("/api/", http.StripPrefix("/api", api))
}

type Server struct {
	Server          *http.Server
	notify          chan error
	Mux             *http.ServeMux
	ShutdownTimeout time.Duration
}

func NewConnectWrapper(s *db.Store, a *auth.Client) *service.ConnectWrapper {
	return service.NewConnectWrapper(s, a)
}

func NewConnectGoServer(
	logger *zerolog.Logger,
	cfg Config,
) *Server {
	mux := http.NewServeMux()
	address := fmt.Sprintf("%s:%d", cfg.ConnectConfig.Host, cfg.ConnectConfig.Port)

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

func (s *Server) start(logger *zerolog.Logger) {
	go func() {
		logger.Info().Str("address", s.Server.Addr).Str("path", "/").Msg("Serving static files")
		logger.Info().Str("address", s.Server.Addr).Str("path", "/api").Msg("Listening for connect-go")
		s.notify <- s.Server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.ShutdownTimeout)
	defer cancel()

	return s.Server.Shutdown(ctx)
}
