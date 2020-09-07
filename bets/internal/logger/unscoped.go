package logger

import (
	"github.com/rs/zerolog"
)

type unscoped struct {
	zerolog.Logger
}

func (u *unscoped) Info(msg string) {
	u.Logger.Info().Caller(1).Msg(msg)
}

func (u *unscoped) Infof(format string, v ...interface{}) {
	u.Logger.Info().Caller(1).Msgf(format, v...)
}

func (u *unscoped) Debug(msg string) {
	u.Logger.Debug().Caller(1).Msg(msg)
}

func (u *unscoped) Debugf(format string, v ...interface{}) {
	u.Logger.Debug().Caller(1).Msgf(format, v...)
}

func (u *unscoped) Error(err error) {
	u.Logger.Error().Caller(1).Msgf("%+v", err)
}

func (u *unscoped) Fatal(err error) {
	u.Logger.Fatal().Caller(1).Msgf("%+v", err)
}
