package logger

import (
	"github.com/rs/zerolog"
)

type httpScoped struct {
	zerolog.Logger
	trace *trace
}

func (h *httpScoped) Info(msg string) {
	h.Logger.Info().Caller(1).Msg(h.trace.String() + msg)
}

func (h *httpScoped) Infof(format string, v ...interface{}) {
	h.Logger.Info().Caller(1).Msgf(h.trace.String()+format, v...)
}

func (h *httpScoped) Debug(msg string) {
	h.Logger.Debug().Caller(1).Msg(h.trace.String() + msg)
}

func (h *httpScoped) Debugf(format string, v ...interface{}) {
	h.Logger.Debug().Caller(1).Msgf(h.trace.String()+format, v...)
}

func (h *httpScoped) Error(err error) {
	h.Logger.Error().Caller(1).Msgf(h.trace.String()+"%+v", err)
}

func (h *httpScoped) Fatal(err error) {
	h.Logger.Fatal().Caller(1).Msgf(h.trace.String()+"%+v", err)
}
