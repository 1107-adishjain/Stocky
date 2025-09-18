package handlers

import (
	"errors"
	"net/http"
	"stocky/internal/database"
	"stocky/internal/models"
	"stocky/internal/services"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
)

func RewardUser(c *gin.Context) {
	var req models.RewardRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.DB
	var existing models.Reward
	if err := db.Where("idempotency_key = ?", req.IdempotencyKey).First(&existing).Error; err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Reward already processed", "reward_id": existing.ID})
		return
	} else if !errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	var user models.User
	db.FirstOrCreate(&user, models.User{UserID: req.UserID})

	reward := models.Reward{
		UserID:         req.UserID,
		StockSymbol:    req.StockSymbol,
		Quantity:       decimal.NewFromFloat(req.Quantity),
		RewardedAt:     time.Now().UTC(),
		IdempotencyKey: req.IdempotencyKey,
	}

	if err := db.Create(&reward).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create reward"})
		return
	}
	
	// Hypothetical prices and fees for ledger entry
	prices, _ := services.GetCurrentPrices([]string{reward.StockSymbol})
	go services.RecordRewardTransaction(reward, prices[reward.StockSymbol], 15.50) // 15.50 INR fee example

	resp := models.RewardResponse{
		ID:          reward.ID,
		UserID:      reward.UserID,
		StockSymbol: reward.StockSymbol,
		Quantity:    req.Quantity,
		RewardedAt:  reward.RewardedAt,
	}

	c.JSON(http.StatusCreated, resp)
}
