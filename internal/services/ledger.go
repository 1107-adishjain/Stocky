package services

import (
	"stocky/internal/database"
	"stocky/internal/models"
	"github.com/shopspring/decimal"
)

func RecordRewardTransaction(reward models.Reward, price float64, fees float64) error {
	db := database.DB
	txID := reward.IdempotencyKey

	priceDecimal := decimal.NewFromFloat(price)
	feesDecimal := decimal.NewFromFloat(fees)
	quantityDecimal := reward.Quantity

	totalCost := quantityDecimal.Mul(priceDecimal).Add(feesDecimal)

	// Simplified: Fetch account IDs. A real system would cache this.
	var cashAccount, equityAccount, feeAccount models.Account
	db.First(&cashAccount, "name = ?", "COMPANY_CASH")
	db.First(&equityAccount, "name = ?", "USER_STOCK_EQUITY")
	db.First(&feeAccount, "name = ?", "FEE_BROKERAGE_EXPENSE") // Example fee account

	entries := []models.LedgerEntry{
		{TransactionID: txID, AccountID: equityAccount.ID, Credit: totalCost},
		{TransactionID: txID, AccountID: cashAccount.ID, Debit: totalCost},
	}

	return db.Create(&entries).Error
}
