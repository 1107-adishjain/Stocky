package handlers

import (
	"net/http"
	"stocky/internal/database"
	"stocky/internal/models"
	"stocky/internal/services"
	
	"github.com/gin-gonic/gin" 
	"github.com/shopspring/decimal"
)

type holding struct {
	StockSymbol   string
	TotalQuantity decimal.Decimal
}

func GetUserPortfolio(c *gin.Context) {
	userID := c.Param("userId")
	db := database.DB

	var holdings []holding
	db.Model(&models.Reward{}).
		Select("stock_symbol, sum(quantity) as total_quantity").
		Where("user_id = ?", userID).
		Group("stock_symbol").
		Scan(&holdings)

	if len(holdings) == 0 {
		c.JSON(http.StatusOK, []models.PortfolioItem{})
		return
	}
	
	symbols := make([]string, len(holdings))
	for i, h := range holdings {
		symbols[i] = h.StockSymbol
	}

	prices, err := services.GetCurrentPrices(symbols)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not fetch prices"})
		return
	}

	resp := make([]models.PortfolioItem, len(holdings))
	for i, h := range holdings {
		qty, _ := h.TotalQuantity.Float64()
		price := prices[h.StockSymbol]
		resp[i] = models.PortfolioItem{
			StockSymbol:   h.StockSymbol,
			TotalQuantity: qty,
			CurrentValue:  price * qty,
		}
	}

	c.JSON(http.StatusOK, resp)
}

