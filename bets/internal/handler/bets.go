package handler

import (
	"context"
	"fmt"
	"log/slog"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/angelokurtis/go-otel/span"
	"github.com/gin-gonic/gin"
	"github.com/lmittmann/tint"

	"github.com/angelokurtis/football-bets/bets/internal/matches"
	"github.com/angelokurtis/football-bets/bets/internal/teams"
)

type Bets struct {
	matchesClient matches.ClientWithResponsesInterface
	teamsClient   teams.ClientWithResponsesInterface
}

func NewBets(matchesClient matches.ClientWithResponsesInterface, teamsClient teams.ClientWithResponsesInterface) *Bets {
	return &Bets{matchesClient: matchesClient, teamsClient: teamsClient}
}

func (s *Bets) Create(c *gin.Context) {
	ctx, end := span.Start(c.Request.Context())
	defer end()

	slog.Info("Create handler started")

	matchRes, err := s.findMatches(ctx)
	if err != nil {
		slog.Error("Error fetching matches", tint.Err(err))
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusServiceUnavailable, err)

		return
	}

	match := s.randomMatch(ctx, matchRes)

	homeTeamID, err := extractTeamId(ctx, match.ScoreHome.Links)
	if err != nil {
		slog.Error("Error extracting home team ID", tint.Err(err))
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	awayTeamID, err := extractTeamId(ctx, match.ScoreAway.Links)
	if err != nil {
		slog.Error("Error extracting away team ID", tint.Err(err))
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	homeTeamRes, err := s.findTeam(ctx, homeTeamID)
	if err != nil {
		slog.Error("Error fetching home team details", tint.Err(err))
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusServiceUnavailable, err)

		return
	}

	awayTeamRes, err := s.findTeam(ctx, awayTeamID)
	if err != nil {
		slog.Error("Error fetching away team details", tint.Err(err))
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusServiceUnavailable, err)

		return
	}

	matchID, err := extractMatchId(ctx, match.Links)
	if err != nil {
		slog.Error("Error extracting match ID", tint.Err(err))
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	championshipRes, err := s.findMatch(ctx, matchID)
	if err != nil {
		slog.Error("Error fetching championship details", tint.Err(err))
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	span.Event(ctx, "Data retrieval successful")
	slog.Info("Data retrieval successful, creating response")

	c.JSON(
		http.StatusOK,
		gin.H{
			"championship": championshipRes.JSON200,
			"date":         match.Date,
			"score_home": gin.H{
				"goals": match.ScoreHome.Goals,
				"bet":   bet(),
				"team":  homeTeamRes.JSON200,
			},
			"score_away": gin.H{
				"goals": match.ScoreAway.Goals,
				"bet":   bet(),
				"team":  awayTeamRes.JSON200,
			},
			"status": match.Status,
			"_links": gin.H{"match": match.Links.Match},
		},
	)
}

func (s *Bets) findMatch(ctx context.Context, matchID string) (*matches.FindChampionshipResponse, error) {
	ctx, end := span.Start(ctx)
	defer end()

	return s.matchesClient.FindChampionshipWithResponse(ctx, matchID)
}

func (s *Bets) findTeam(ctx context.Context, homeTeamID string) (*teams.FindOneResponse, error) {
	ctx, end := span.Start(ctx)
	defer end()

	return s.teamsClient.FindOneWithResponse(ctx, homeTeamID)
}

func (s *Bets) findMatches(ctx context.Context) (*matches.FindAllResponse, error) {
	ctx, end := span.Start(ctx)
	defer end()

	return s.matchesClient.FindAllWithResponse(ctx)
}

func (s *Bets) randomMatch(ctx context.Context, matchRes *matches.FindAllResponse) *matches.Match {
	matchList := *matchRes.JSON200.Embedded.Matches
	index := rand.Intn(len(matchList))
	match := &matchList[index]

	return match
}

func extractTeamId(ctx context.Context, links *matches.ScoreLinks) (string, error) {
	href := *links.Team.Href
	regex := regexp.MustCompile(`/teams/(.+)`)

	found := regex.FindStringSubmatch(href)
	if len(found) != 2 {
		return "", fmt.Errorf("failed to extract team ID from the provided link %q", href)
	}

	return found[1], nil
}

func extractMatchId(ctx context.Context, link *matches.MatchLinks) (string, error) {
	href := *link.Self.Href
	regex := regexp.MustCompile(`/matches/(.+)`)

	found := regex.FindStringSubmatch(href)
	if len(found) != 2 {
		return "", fmt.Errorf("failed to extract match ID from the provided link %q", href)
	}

	return found[1], nil
}

func bet() int {
	minimum := 0
	maximum := 5

	return rand.Intn(maximum-minimum) + minimum
}
