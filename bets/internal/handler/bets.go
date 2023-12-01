package handler

import (
	"fmt"
	"math/rand"
	"net/http"
	"regexp"

	"github.com/angelokurtis/go-otel/span"
	"github.com/gin-gonic/gin"

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

	matchRes, err := s.matchesClient.FindAllWithResponse(ctx)
	if err != nil {
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusServiceUnavailable, err)

		return
	}

	matchList := *matchRes.JSON200.Embedded.Matches
	index := rand.Intn(len(matchList))
	match := &matchList[index]

	homeTeamID, err := extractTeamId(match.ScoreHome.Links)
	if err != nil {
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	awayTeamID, err := extractTeamId(match.ScoreAway.Links)
	if err != nil {
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	homeTeamRes, err := s.teamsClient.FindOneWithResponse(ctx, homeTeamID)
	if err != nil {
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusServiceUnavailable, err)

		return
	}

	awayTeamRes, err := s.teamsClient.FindOneWithResponse(ctx, awayTeamID)
	if err != nil {
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusServiceUnavailable, err)

		return
	}

	matchID, err := extractMatchId(match.Links)
	if err != nil {
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

	championshipRes, err := s.matchesClient.FindChampionshipWithResponse(ctx, matchID)
	if err != nil {
		_ = span.Error(ctx, err)
		_ = c.AbortWithError(http.StatusInternalServerError, err)

		return
	}

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

func extractTeamId(links *matches.ScoreLinks) (string, error) {
	href := *links.Team.Href
	regex := regexp.MustCompile(`/teams/(.+)`)

	found := regex.FindStringSubmatch(href)
	if len(found) != 2 {
		return "", fmt.Errorf("failed to extract team ID from the provided link %q", href)
	}

	return found[1], nil
}

func extractMatchId(link *matches.MatchLinks) (string, error) {
	href := *link.Self.Href
	regex := regexp.MustCompile(`/matches/(.+)`)

	found := regex.FindStringSubmatch(href)
	if len(found) != 2 {
		return "", fmt.Errorf("failed to extract match ID from the provided link %q", href)
	}

	return found[1], nil
}

func bet() int {
	min := 0
	max := 5

	return rand.Intn(max-min) + min
}
