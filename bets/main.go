package main

import (
	healthCheck "github.com/RaMin0/gin-health-check"
	"github.com/angelokurtis/football-bets/bets/internal/logger"
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
		log := logger.New(c)
		log.Info("received request to bet")

		m, err := matches.GetRandomly(c)
		if err != nil || m == nil {
			log.Error(err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "matches service is unavailable"})
			return
		}
		log.Info("selected a match randomly")

		ht, err := teams.GetOne(c, m.ScoreHome.Links.Team.Href)
		if err != nil || ht == nil {
			log.Error(err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "teams service is unavailable"})
			return
		}
		log.Infof("obtained %s as home team", ht.Name)

		at, err := teams.GetOne(c, m.ScoreAway.Links.Team.Href)
		if err != nil || at == nil {
			log.Error(err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "teams service is unavailable"})
			return
		}
		log.Infof("obtained %s as away team", at.Name)

		championship, err := matches.GetChampionship(c, m.Links.Championship.Href)
		if err != nil || championship == nil {
			log.Error(err)
			c.JSON(http.StatusServiceUnavailable, gin.H{"status": "championships service is unavailable"})
			return
		}
		log.Infof("obtained %s championship", championship.Name)

		bet := bets.New(m, championship, ht, at)
		log.Info("successful bet")
		c.JSON(http.StatusCreated, bet)
	})
	err := r.Run(":9090")
	if err != nil {
		panic(err)
	}
}
