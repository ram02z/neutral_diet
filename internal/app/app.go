package app

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/ram02z/neutral_diet/internal/app/connectgo"
	"github.com/ram02z/neutral_diet/internal/app/firebase"
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

	err := godotenv.Load(".env." + env + ".local")
	if err != nil {
		l.Warn().Msg(err.Error())
	}
	if env != "test" {
		err = godotenv.Load(".env.local")
		if err != nil {
			l.Warn().Msg(err.Error())
		}
	}
	err = godotenv.Load(".env." + env)
	if err != nil {
		l.Warn().Msg(err.Error())
	}
	err = godotenv.Load()
	if err != nil {
		l.Warn().Msg(err.Error())
	}

	// Firebase services
	firebaseCfg, err := firebase.NewConfig()
	if err != nil {
		l.Fatal().Err(err).Msg("Could not create firebase config")
	}
	firebaseApp, err := firebase.NewApp(firebaseCfg)
	if err != nil {
		l.Fatal().Err(err)
	}
	authClient, err := firebaseApp.NewAuth()
	if err != nil {
		l.Fatal().Err(err)
	}
	l.Info().Msg("Successfully initialised Firebase Auth client")
	messagingClient, err := firebaseApp.NewMessaging()
	if err != nil {
		l.Fatal().Err(err)
	}
	l.Info().Msg("Successfully initialised Firebase Messaging client")

	// Database service
	pgpool, err := sql.NewDatabase()
	if err != nil {
		l.Fatal().Err(err).Msg("Could not connect to postgres")
	}
	err = pgpool.Ping(context.Background())
	if err != nil {
		l.Fatal().Err(err).Msg("Could not ping the DB")
	}
	l.Info().Msg("Successfully connected to postgres")
	defer func() {
		l.Info().Msg("Closing DB connection")
		pgpool.Close()
	}()

	dataStore := service.NewDataStore(pgpool)

	// Cron service
	cronScheduler := gocron.NewScheduler(time.UTC)
	jobContext := l.WithContext(context.Background())
	jobWrapper := service.NewJobWrapper(pgpool, messagingClient, &jobContext)
	for _, job := range jobWrapper.Jobs() {
		_, err := cronScheduler.
			Every(1).
			Day().
			At(job.At).
			Do(job.Job)
		if err != nil {
			l.Err(err).Msg("Failed to add job")
		}
	}
	cronScheduler.StartAsync()
	l.Info().Msg("Successfully started the cron scheduler")

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
		AuthClient: authClient,
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
