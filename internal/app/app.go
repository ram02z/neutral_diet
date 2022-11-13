package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/ram02z/neutral_diet/internal/app/connectgo"
	"github.com/ram02z/neutral_diet/internal/app/logging"
	"github.com/ram02z/neutral_diet/internal/app/service"
	"github.com/ram02z/neutral_diet/internal/app/sql"
)

func Run() {
	l := logging.NewLogger()

	// Database service
	pgpool, err := sql.NewDatabase()
	if err != nil {
		l.Fatal().Err(err).Msg("Could not connect to postgres")
	}
	l.Info().Msg("Successfully connected to postgres")
	defer func() {
		l.Info().Msg("Closing DB connection")
		pgpool.Close()
	}()

	dataStore := service.NewDataStore(pgpool)
	service := service.NewService(dataStore)

	// Connect service
	connectCfg, err := connectgo.NewConfig()
	if err != nil {
		l.Fatal().Err(err).Msg("Could not create connect-go config")
	}
	connectWrapper := connectgo.NewConnectWrapper(service)
	server := connectgo.NewConnectGoServer(l, connectCfg)
	registerIn := connectgo.RegisterConnectGoServerInput{
		Logger:     l,
		ConnectSvc: connectWrapper,
		Mux:        server.Mux,
	}
	connectgo.RegisterConnectGoServer(registerIn)

	l.Info().Msg("App - running")

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info().Msg("App - signal: " + s.String())
	case err = <-server.Notify():
		l.Error().Err(err).Msg("App - server.Notify")
	}

	// Shutdown
	err = server.Shutdown()
	if err != nil {
		l.Error().Err(err).Msg("App - server.Shutdown")
	}
	l.Info().Msg("App - server.Shutdown")
}
