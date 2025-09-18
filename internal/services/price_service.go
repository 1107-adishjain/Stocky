package services

import (
	"math/rand"
	"time"
)

var prices = make(map[string]float64)
var lastUpdated = make(map[string]time.Time)

func GetCurrentPrices(symbols []string) (map[string]float64, error) {
	rand.Seed(time.Now().UnixNano())
	results := make(map[string]float64)

	for _, symbol := range symbols {
		if time.Since(lastUpdated[symbol]) > time.Hour {
			var price float64
			switch symbol {
			case "RELIANCE":
				price = 2800 + rand.Float64()*(3200-2800)
			case "TCS":
				price = 3500 + rand.Float64()*(4000-3500)
			case "INFY":
				price = 1400 + rand.Float64()*(1600-1400)
			default:
				price = 100 + rand.Float64()*(5000-100)
			}
			prices[symbol] = price
			lastUpdated[symbol] = time.Now()
		}
		results[symbol] = prices[symbol]
	}

	return results, nil
}
