package app

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/ram02z/neutral_diet/internal/app/auth"
	"github.com/ram02z/neutral_diet/internal/app/connectgo"
	"github.com/ram02z/neutral_diet/internal/app/logging"
	"github.com/ram02z/neutral_diet/internal/app/service"
	"github.com/ram02z/neutral_diet/internal/app/sql"
)

func Run() {
	l := logging.NewLogger()

	// Load .env variables
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if env != "test" {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load()

	// Firebase Auth service
	firebaseCfg, err := auth.NewConfig()
	if err != nil {
		l.Fatal().Err(err).Msg("Could not create firebase config")
	}
	authClient, err := auth.NewAuth(firebaseCfg)
	if err != nil {
		l.Fatal().Err(err)
	}

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

	// Connect service
	connectCfg, err := connectgo.NewConfig()
	if err != nil {
		l.Fatal().Err(err).Msg("Could not create connect-go config")
	}
	connectWrapper := connectgo.NewConnectWrapper(dataStore, authClient)
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
