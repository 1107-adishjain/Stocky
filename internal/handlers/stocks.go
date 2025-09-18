package handlers

import (
	"net/http"
	"stocky/internal/database"
	"stocky/internal/models"
	"time"

	"github.com/gin-gonic/gin"
)

func GetTodayStocks(c *gin.Context) {
	userID := c.Param("userId")
	db := database.DB

	todayStart := time.Now().UTC().Truncate(24 * time.Hour)
	var rewards []models.Reward
	db.Where("user_id = ? AND rewarded_at >= ?", userID, todayStart).Find(&rewards)

	resp := make([]models.DailyStockRecord, len(rewards))
	for i, r := range rewards {
		qty, _ := r.Quantity.Float64()
		resp[i] = models.DailyStockRecord{
			StockSymbol: r.StockSymbol,
			Quantity:    qty,
			RewardedAt:  r.RewardedAt,
		}
	}
	c.JSON(http.StatusOK, resp)
}

func GetHistoricalINR(c *gin.Context) {
	// A production implementation would query pre-aggregated daily snapshots.
	// This is a placeholder for demonstration.
	resp := []models.HistoricalValue{
		{Date: "2025-09-18", INRValue: 12500.75},
		{Date: "2025-09-17", INRValue: 11800.20},
	}
	c.JSON(http.StatusOK, resp)
}
