package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type User struct {
	ID        uint   `gorm:"primaryKey"`
	UserID    string `gorm:"uniqueIndex;not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Reward struct {
	ID             uint            `gorm:"primaryKey"`
	UserID         string          `gorm:"index;not null"`
	StockSymbol    string          `gorm:"index;not null"`
	Quantity       decimal.Decimal `gorm:"type:numeric(18,6);not null"`
	RewardedAt     time.Time       `gorm:"not null"`
	IdempotencyKey string          `gorm:"uniqueIndex;not null"`
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Account struct {
	ID        uint   `gorm:"primaryKey"`
	Name      string `gorm:"unique;not null"`
	Type      string `gorm:"not null"` 
	CreatedAt time.Time
	UpdatedAt time.Time
}

type LedgerEntry struct {
	ID            uint            `gorm:"primaryKey"`
	TransactionID string          `gorm:"index;not null"`
	AccountID     uint            `gorm:"not null"`
	Account       Account         `gorm:"foreignKey:AccountID"`
	Debit         decimal.Decimal `gorm:"type:numeric(18,4)"` 
	Credit        decimal.Decimal `gorm:"type:numeric(18,4)"` 
	StockSymbol   *string         `gorm:"index"`
	EntryAt       time.Time       `gorm:"not null"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// DTOs for API requests and responses
type RewardRequest struct {
	UserID         string  `json:"user_id" binding:"required"`
	StockSymbol    string  `json:"stock_symbol" binding:"required"`
	Quantity       float64 `json:"quantity" binding:"required"`
	IdempotencyKey string  `json:"idempotency_key" binding:"required"`
}

type RewardResponse struct {
	ID          uint      `json:"id"`
	UserID      string    `json:"user_id"`
	StockSymbol string    `json:"stock_symbol"`
	Quantity    float64   `json:"quantity"`
	RewardedAt  time.Time `json:"rewarded_at"`
}

type DailyStockRecord struct {
	StockSymbol string    `json:"stock_symbol"`
	Quantity    float64   `json:"quantity"`
	RewardedAt  time.Time `json:"rewarded_at"`
}

type HistoricalValue struct {
	Date     string  `json:"date"`
	INRValue float64 `json:"inr_value"`
}

type PortfolioItem struct {
	StockSymbol   string  `json:"stock_symbol"`
	TotalQuantity float64 `json:"total_quantity"`
	CurrentValue  float64 `json:"current_value_inr"`
}

type StatsResponse struct {
	TotalSharesToday      map[string]float64 `json:"total_shares_today"`
	CurrentPortfolioValue float64            `json:"current_portfolio_value_inr"`
}
