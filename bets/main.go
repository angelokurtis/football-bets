package main

import (
	healthCheck "github.com/RaMin0/gin-health-check"
	"github.com/angelokurtis/football-bets/bets/internal/log"
	"github.com/angelokurtis/football-bets/bets/pkg/bets"
	"github.com/angelokurtis/football-bets/bets/pkg/matches"
	"github.com/angelokurtis/football-bets/bets/pkg/teams"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(healthCheck.Default())
	r.POST("/bets", func(c *gin.Context) {
		log.Info("received request to bet")
		headers := c.Request.Header.Clone()
		m, err := matches.GetRandomly(headers)
		if err != nil || m == nil {
			log.Error(err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "matches service is unavailable"})
			return
		}
		log.Info("selected a match randomly")

		ht, err := teams.GetOne(m.ScoreHome.Links.Team.Href, headers)
		if err != nil || ht == nil {
			log.Error(err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "teams service is unavailable"})
			return
		}
		log.Info("obtained home team")

		at, err := teams.GetOne(m.ScoreAway.Links.Team.Href, headers)
		if err != nil || at == nil {
			log.Error(err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "teams service is unavailable"})
			return
		}
		log.Info("obtained away team")

		championship, err := matches.GetChampionship(m.Links.Championship.Href, headers)
		if err != nil || championship == nil {
			log.Error(err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "championships service is unavailable"})
			return
		}
		log.Info("obtained championship")

		bet := bets.New(m, championship, ht, at)
		log.Info("successful bet")
		c.JSON(http.StatusCreated, bet)
	})
	err := r.Run(":9090")
	if err != nil {
		panic(err)
	}
}
