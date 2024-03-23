package service

import (
	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/pkg/utils"
)

type ProcessSellOutput struct {
	Profit    float64
	CGTProfit float64
	BuysSold  []repository.BuySold
}

// Calculate the profit on the sell using FIFO
// buy.QuantityLeft in the relevant buys of the input buy array is modified to reflect the sale
func fifo(sell repository.Transaction, buys *[]repository.Buy) ProcessSellOutput {
	profit, cgtProfit := 0.0, 0.0
	// List of buys that the sell corresponds to
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

			buysSold = append(buysSold, repository.BuySold{
				BuyID:    buy.ID,
				Quantity: sellQ,
			})
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

	return ProcessSellOutput{
		Profit:    profit,
		CGTProfit: cgtProfit,
		BuysSold:  buysSold,
	}
}
