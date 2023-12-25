package service

import (
	"context"

	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/pkg/utils"
)

type FifoOutput struct {
	Profit    float64
	CGTProfit float64
	BuysSold  []repository.BuySold
}

// Calculate the profit on the trade using fifo
func fifo(ctx context.Context, sell repository.Transaction, buys *[]repository.Buy) FifoOutput {
	profit, cgtProfit := 0.0, 0.0
	buysSold := []repository.BuySold{}

	sellQ := sell.Quantity
	sellPrice := sell.RealPrice

	// since buys are in FIFO order, we can just iterate through them
	for i := 0; i < len(*buys) && sellQ > 0; i++ {
		buy := (*buys)[i]
		buyQ := buy.QuantityLeft
		// If current lot sold continue
		if buyQ == 0 {
			continue
		}

		tradeProfit := 0.0
		buyPrice := buy.RealPrice
		// If sell units >= buy units
		if sellQ >= buyQ {
			tradeProfit = buyQ * (sellPrice - buyPrice)
			sellQ -= buyQ

			buysSold = append(buysSold, repository.BuySold{
				BuyID:    buy.ID,
				Quantity: buyQ,
			})
			buy.QuantityLeft = 0
		} else {
			tradeProfit = sellQ * (sellPrice - buyPrice)
			buy.QuantityLeft -= sellQ
			sellQ = 0
		}
		profit += tradeProfit

		// Calculate cgt profit
		cgtDiscount := utils.IsOneYearGreaterThan(buy.Date, sell.Date)
		if tradeProfit > 0 && cgtDiscount {
			cgtProfit += tradeProfit / 2
		} else {
			cgtProfit += tradeProfit
		}

		// We need to update the buy as we've changed the quantity left
		(*buys)[i] = buy

	}

	return FifoOutput{
		Profit:    profit,
		CGTProfit: cgtProfit,
		BuysSold:  buysSold,
	}
}
