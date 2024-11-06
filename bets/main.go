//go:generate go run -mod=mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 -generate std-http,spec -package bets -o internal/bets/server.go docs/bets.yaml
//go:generate go run -mod=mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 -generate client -package matches -o internal/matches/client.go docs/matches.yaml
//go:generate go run -mod=mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 -generate types -package matches -o internal/matches/types.go docs/matches.yaml
//go:generate go run -mod=mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 -generate client -package teams -o internal/teams/client.go docs/teams.yaml
//go:generate go run -mod=mod github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@v2.4.1 -generate types -package teams -o internal/teams/types.go docs/teams.yaml

package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/angelokurtis/go-otel/starter"
	"github.com/lmittmann/tint"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/angelokurtis/football-bets/bets/internal/bets"
	"github.com/angelokurtis/football-bets/bets/internal/handler"
	"github.com/angelokurtis/football-bets/bets/internal/httpclient"
	"github.com/angelokurtis/football-bets/bets/internal/matches"
	"github.com/angelokurtis/football-bets/bets/internal/otel"
	"github.com/angelokurtis/football-bets/bets/internal/teams"
)

func init() {
	logger := slog.New(otel.NewLogHandler(tint.NewHandler(os.Stderr, &tint.Options{
		AddSource:  true,
		Level:      slog.LevelDebug,
		TimeFormat: time.Kitchen,
	})))
	slog.SetDefault(logger)
}

func main() {
	ctx := context.Background()

	_, shutdown, err := starter.StartProviders(ctx)
	defer shutdown()

	if err != nil {
		slog.ErrorContext(ctx, "Error starting OpenTelemetry providers", tint.Err(err))
		return
	}

	slog.InfoContext(ctx, "Starting application...")

	httpClient := httpclient.New()

	matchesClient, err := matches.NewClientWithHTTPClient(httpClient)
	if err != nil {
		slog.ErrorContext(ctx, "Error creating matches client", tint.Err(err))
		return
	}

	teamsClient, err := teams.NewClientWithHTTPClient(httpClient)
	if err != nil {
		slog.ErrorContext(ctx, "Error creating teams client", tint.Err(err))
		return
	}

	h := otelhttp.NewHandler(
		bets.HandlerFromMux(
			handler.NewBets(matchesClient, teamsClient),
			http.NewServeMux(),
		),
		"",
		otelhttp.WithSpanNameFormatter(otel.FormatSpanName),
	)

	addr := ":8081"
	slog.InfoContext(ctx, "Starting server", slog.String("addr", addr))

	if err = (&http.Server{
		Addr:    addr,
		Handler: h,
	}).ListenAndServe(); err != nil {
		slog.ErrorContext(ctx, "Error starting HTTP server", tint.Err(err))
		return
	}

	slog.InfoContext(ctx, "Application stopped")
}
