package bets

import (
	"github.com/angelokurtis/football-bets/bets/pkg/championships"
	"github.com/angelokurtis/football-bets/bets/pkg/matches"
	"github.com/angelokurtis/football-bets/bets/pkg/teams"
	"math/rand"
	"time"
)

type (
	Bet struct {
		Links        *Links `json:"_links"`
		Date         string `json:"date"`
		ScoreHome    *Score `json:"score_home"`
		ScoreAway    *Score `json:"score_away"`
		Championship *championships.Championship
		Status       string `json:"status"`
	}
	Score struct {
		Team  *teams.Team `json:"team"`
		Goals int         `json:"goals"`
		Bet   int         `json:"bet"`
	}
	Link struct {
		Href string `json:"href"`
	}
	Links struct {
		Match *Link `json:"match"`
	}
)

func New(m *matches.Match, c *championships.Championship, ht *teams.Team, at *teams.Team) *Bet {
	return &Bet{
		Links: &Links{Match: &Link{Href: m.Links.Self.Href}},
		Date:  m.Date,
		ScoreHome: &Score{
			Team:  ht,
			Goals: m.ScoreHome.Goals,
			Bet:   bet(),
		},
		ScoreAway: &Score{
			Team:  at,
			Goals: m.ScoreAway.Goals,
			Bet:   bet(),
		},
		Championship: c,
		Status:       m.Status,
	}
}

func bet() int {
	rand.Seed(time.Now().UnixNano())
	min := 0
	max := 5
	return rand.Intn(max-min+1) + min
}
