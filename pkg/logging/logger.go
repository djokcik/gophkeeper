package logging

import (
	"github.com/caarlos0/env/v6"
	"github.com/rs/zerolog"
	"os"
	"time"
)

var cfg loggerConfig

type loggerConfig struct {
	LogPath string `env:"LOG_PATH" envDefault:"/tmp/log.txt"`
}

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano
	zerolog.TimestampFieldName = "@timestamp"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "message"

	err := env.Parse(&cfg)
	if err != nil {
		NewLogger().Fatal().Err(err).Msg("loggerConfig: error parse environment")
	}
}

func NewFileLogger() *zerolog.Logger {
	file, err := os.Create(cfg.LogPath)
	if err != nil {
		NewLogger().Fatal().Err(err).Msg("invalid open log file")
	}

	logger := zerolog.New(os.Stdout).
		Output(file).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()

	return &logger
}

// NewLogger creates a new customizable logger.
func NewLogger() *zerolog.Logger {
	logger := zerolog.New(os.Stdout).
		Output(zerolog.ConsoleWriter{Out: os.Stdout}).
		Level(zerolog.TraceLevel).
		With().
		Timestamp().
		Logger()

	return &logger
}
