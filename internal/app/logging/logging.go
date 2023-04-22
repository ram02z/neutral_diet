// Package logging implements logging utilities.
package logging

import (
	"os"

	"github.com/rs/zerolog"
)

// NewLogger returns instance of logger with timestamp and caller fields.
func NewLogger() *zerolog.Logger {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger := zerolog.New(os.Stderr).With().Timestamp().Caller().Logger()

	return &logger
}
