//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.0 -generate gin,spec -package bets -o internal/bets/server.go docs/bets.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.0 -generate client -package matches -o internal/matches/client.go docs/matches.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.0 -generate types -package matches -o internal/matches/types.go docs/matches.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.0 -generate client -package teams -o internal/teams/client.go docs/teams.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/cmd/oapi-codegen@v1.13.0 -generate types -package teams -o internal/teams/types.go docs/teams.yaml

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/angelokurtis/football-bets/bets/internal/bets"
	"github.com/angelokurtis/football-bets/bets/internal/handler"
	"github.com/angelokurtis/football-bets/bets/internal/httpclient"
	"github.com/angelokurtis/football-bets/bets/internal/matches"
	"github.com/angelokurtis/football-bets/bets/internal/otel"
	"github.com/angelokurtis/football-bets/bets/internal/teams"
)

func main() {
	ctx := context.Background()

	tp, err := otel.Init(ctx)
	if err != nil {
		log.Fatal(err)
	}

	defer func() { _ = tp.Shutdown(ctx) }()

	var (
		router     = gin.Default()
		httpClient = httpclient.New(otel.NewTransport())
	)

	matchesClient, err := matches.NewClientWithHTTPClient(httpClient)
	if err != nil {
		log.Fatal(err)
	}

	teamsClient, err := teams.NewClientWithHTTPClient(httpClient)
	if err != nil {
		log.Fatal(err)
	}

	bets.RegisterHandlersWithOptions(router, handler.NewBets(matchesClient, teamsClient), bets.GinServerOptions{BaseURL: ""})
	log.Fatal((&http.Server{
		Addr:    ":8081",
		Handler: otelhttp.NewHandler(router, "", otelhttp.WithSpanNameFormatter(otel.SpanFormatter)),
	}).ListenAndServe())
}
