package handlers

import (
	"net/http"
	"stocky/internal/database"
	"stocky/internal/models"
	"stocky/internal/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
)

func GetUserStats(c *gin.Context) {
	userID := c.Param("userId")
	db := database.DB

	todayStart := time.Now().UTC().Truncate(24 * time.Hour)
	var todayRewards []models.Reward
	db.Where("user_id = ? AND rewarded_at >= ?", userID, todayStart).Find(&todayRewards)
	
	sharesToday := make(map[string]float64)
	for _, r := range todayRewards {
		qty, _ := r.Quantity.Float64()
		sharesToday[r.StockSymbol] += qty
	}

	// Get current portfolio value
	type holding struct {
		StockSymbol string
		TotalQuantity decimal.Decimal
	}
	var holdings []holding
	db.Model(&models.Reward{}).
		Select("stock_symbol, sum(quantity) as total_quantity").
		Where("user_id = ?", userID).
		Group("stock_symbol").
		Scan(&holdings)
	
	var totalValue float64
	if len(holdings) > 0 {
		symbols := make([]string, len(holdings))
		for i, h := range holdings {
			symbols[i] = h.StockSymbol
		}
		prices, _ := services.GetCurrentPrices(symbols)
		for _, h := range holdings {
			qty, _ := h.TotalQuantity.Float64()
			totalValue += qty * prices[h.StockSymbol]
		}
	}
	
	c.JSON(http.StatusOK, models.StatsResponse{
		TotalSharesToday:      sharesToday,
		CurrentPortfolioValue: totalValue,
	})
}
