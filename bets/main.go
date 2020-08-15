package main

import (
	"github.com/angelokurtis/football-bets/bets/pkg/bets"
	"github.com/angelokurtis/football-bets/bets/pkg/championships"
	"github.com/angelokurtis/football-bets/bets/pkg/matches"
	"github.com/angelokurtis/football-bets/bets/pkg/teams"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/bets", func(c *gin.Context) {
		m, err := matches.GetRandomly()
		if err != nil || m == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "matches service is unavailable"})
			return
		}
		ht, err := teams.GetOne(m.ScoreHome.Links.Team.Href)
		if err != nil || ht == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "teams service is unavailable"})
			return
		}
		at, err := teams.GetOne(m.ScoreAway.Links.Team.Href)
		if err != nil || at == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "teams service is unavailable"})
			return
		}
		championship, err := championships.GetOne(m.Links.Championship.Href)
		if err != nil || championship == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "championships service is unavailable"})
			return
		}
		bet := bets.New(m, championship, ht, at)
		c.JSON(http.StatusCreated, bet)
	})
	err := r.Run(":9090")
	if err != nil {
		panic(err)
	}
}
