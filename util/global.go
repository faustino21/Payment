package util

import (
	"github.com/rs/zerolog"
	"os"
)

var Log zerolog.Logger

func NewLog(loglevel string) {
	zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	if loglevel == "debug" {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()
	Log = logger
}

func LogError(funcName, position string, err error) {
	Log.Error().Msgf(funcName+"%s : %v", position, err)
}
