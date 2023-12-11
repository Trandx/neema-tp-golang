package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Logger = NewZeroLogger()

type ZeroLogger struct {
	logger zerolog.Logger
}

func NewZeroLogger() *ZeroLogger {
	output := zerolog.ConsoleWriter{Out: os.Stdout}
	output.TimeFormat = "2006-01-02 15:04:05"
	logger := zerolog.New(output).With().Timestamp().Logger()

	zerolog.SetGlobalLevel(zerolog.InfoLevel)

	return &ZeroLogger{logger: logger}
}

// Info logs an informational message.
func (l *ZeroLogger) Info(message string) {
	l.logger.Info().Msg(message)
}

// Error logs an error message.
func (l *ZeroLogger) Error(message string) {
	l.logger.Error().Msg(message)
}

// Debug logs a debug message.
func (l *ZeroLogger) Debug(message string) {
	l.logger.Debug().Msg(message)
}

// Panic logs a panic message and panics.
func (l *ZeroLogger) Panic(message string) {
	l.logger.Panic().Msg(message)
}
