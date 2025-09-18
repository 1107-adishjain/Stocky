package api

import (
	"stocky/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Welcome to the Stocky API"})
	})

	r.POST("/reward", handlers.RewardUser)
	r.GET("/today-stocks/:userId", handlers.GetTodayStocks)
	r.GET("/historical-inr/:userId", handlers.GetHistoricalINR)
	r.GET("/stats/:userId", handlers.GetUserStats)
	r.GET("/portfolio/:userId", handlers.GetUserPortfolio)

	return r
}
