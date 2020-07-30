package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(gin.Recovery())
	r.POST("/api/bets", func(c *gin.Context) {
		c.JSON(http.StatusCreated, gin.H{"status": "your bet has been registered"})
	})
	err := r.Run()
	if err != nil {
		panic(err)
	}
}
