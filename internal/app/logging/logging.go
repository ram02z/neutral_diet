package logging

import (
	"os"

	"github.com/rs/zerolog"
)

func NewLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()

	return &logger
}
