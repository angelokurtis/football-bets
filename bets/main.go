//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0 -generate gin,spec -package bets -o internal/bets/server.go docs/bets.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0 -generate client -package matches -o internal/matches/client.go docs/matches.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0 -generate types -package matches -o internal/matches/types.go docs/matches.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0 -generate client -package teams -o internal/teams/client.go docs/teams.yaml
//go:generate go run -mod=mod github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@v2.0.0 -generate types -package teams -o internal/teams/types.go docs/teams.yaml

package main

import (
	"context"
	"log"
	"net/http"

	"github.com/angelokurtis/go-otel/starter"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"

	"github.com/angelokurtis/football-bets/bets/internal/bets"
	"github.com/angelokurtis/football-bets/bets/internal/handler"
	"github.com/angelokurtis/football-bets/bets/internal/httpclient"
	"github.com/angelokurtis/football-bets/bets/internal/matches"
	"github.com/angelokurtis/football-bets/bets/internal/teams"
)

func main() {
	ctx := context.Background()

	_, shutdown, err := starter.StartProviders(ctx)
	defer shutdown()

	if err != nil {
		log.Fatal(err)
	}

	var (
		router     = gin.Default()
		httpClient = httpclient.New(otelhttp.NewTransport(http.DefaultTransport))
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
		Addr: ":8081",
		Handler: otelhttp.NewHandler(router, "", otelhttp.WithSpanNameFormatter(func(_ string, r *http.Request) string {
			return r.Method + " " + r.URL.Path
		})),
	}).ListenAndServe())
}
