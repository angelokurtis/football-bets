//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0 -generate gin,spec -package bets -o internal/bets/server.go docs/bets.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0 -generate client -package matches -o internal/matches/client.go docs/matches.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0 -generate types -package matches -o internal/matches/types.go docs/matches.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0 -generate client -package teams -o internal/teams/client.go docs/teams.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0 -generate types -package teams -o internal/teams/types.go docs/teams.yaml

package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/angelokurtis/go-otel/span"
	"github.com/angelokurtis/go-otel/starter"
	"github.com/gin-gonic/gin"
	"github.com/lmittmann/tint"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/angelokurtis/football-bets/bets/internal/bets"
	"github.com/angelokurtis/football-bets/bets/internal/handler"
	"github.com/angelokurtis/football-bets/bets/internal/httpclient"
	"github.com/angelokurtis/football-bets/bets/internal/matches"
	"github.com/angelokurtis/football-bets/bets/internal/teams"
)

func init() {
	logger := slog.New(tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  true,
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	}))
	slog.SetDefault(logger)
}

func main() {
	ctx := context.Background()

	_, shutdown, err := starter.StartProviders(ctx)
	defer shutdown()

	if err != nil {
		slog.Error("Error starting OpenTelemetry providers", tint.Err(err))
		return
	}

	slog.Info("Starting application...")

	var (
		router     = gin.Default()
		httpClient = httpclient.New(otelhttp.NewTransport(http.DefaultTransport))
	)

	matchesClient, err := matches.NewClientWithHTTPClient(httpClient)
	if err != nil {
		slog.Error("Error creating matches client", tint.Err(err))
		return
	}

	teamsClient, err := teams.NewClientWithHTTPClient(httpClient)
	if err != nil {
		slog.Error("Error creating teams client", tint.Err(err))
		return
	}

	router.Use(func(c *gin.Context) {
		req := c.Request
		reqCtx := otel.GetTextMapPropagator().Extract(req.Context(), propagation.HeaderCarrier(c.Request.Header))

		reqCtx, end := span.StartWithName(reqCtx, req.Method+" "+c.FullPath(),
			oteltrace.WithSpanKind(oteltrace.SpanKindServer),
		)
		defer end()

		// pass the span through the request context
		c.Request = c.Request.WithContext(reqCtx)

		// serve the request to the next middleware
		c.Next()

		span.Attributes(reqCtx,
			semconv.HTTPMethod(req.Method),
			semconv.HTTPRoute(c.FullPath()),
			semconv.HTTPScheme("http"),
			semconv.HTTPStatusCode(c.Writer.Status()),
			semconv.HTTPTarget(req.URL.Path),
		)
	})

	bets.RegisterHandlersWithOptions(router, handler.NewBets(matchesClient, teamsClient), bets.GinServerOptions{BaseURL: ""})

	addr := ":8081"
	slog.Info("Starting server", slog.String("addr", addr))

	if err = (&http.Server{
		Addr:    addr,
		Handler: router,
	}).ListenAndServe(); err != nil {
		slog.Error("Error starting HTTP server", tint.Err(err))
		return
	}

	slog.Info("Application stopped")
}
