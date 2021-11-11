package logger

import (
	"github.com/rs/zerolog"
	"os"
	"time"
)

var TraceId string

func NewLogger() *zerolog.Logger {
	output := zerolog.ConsoleWriter{
		Out:        os.Stdout,
		NoColor:    true,
		TimeFormat: "2006-01-02 15:04:05.000",
		FormatMessage: func(i interface{}) string {
			return "traceId=" + TraceId + " " + i.(string)
		},
	}

	zerolog.TimeFieldFormat = time.RFC3339Nano
	log := zerolog.New(output).With().Timestamp().Logger()

	return &log
}

func Error() *zerolog.Event {

	return NewLogger().Error().Str("app", "online-store")
}

func Info() *zerolog.Event {

	return NewLogger().Info().Str("app", "online-store")
}
