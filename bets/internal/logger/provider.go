package logger

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"time"
)

func New(c context.Context) Logger {
	logger := log.Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.Stamp})
	if g, ok := c.(*gin.Context); ok {
		return &httpScoped{Logger: logger, trace: traceFromHeader(g.Request.Header)}
	} else {
		return &unscoped{Logger: logger}
	}
}
