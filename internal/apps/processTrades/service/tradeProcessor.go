package service

import (
	"context"
	"fmt"

	"github.com/Arjun-P17/tax-go/internal/models"
	"github.com/Arjun-P17/tax-go/internal/utils"
	pkgutils "github.com/Arjun-P17/tax-go/pkg/utils"
)

func (s *Service) ProcessTrades(ctx context.Context) error {
	stocksTransactions, err := s.dbConnector.GetAllStockTransactions(ctx)
	if err != nil {
		return err
	}

	for _, stockTransaction := range stocksTransactions {
		ticker := stockTransaction.Ticker

		stockPosition, err := s.dbConnector.GetStockPositionOrDefault(ctx, ticker)
		if err != nil {
			return err
		}

		for _, transaction := range stockTransaction.Transactions {
			if !utils.IsUniqueTransaction(stockPosition.Buys, transaction) && !utils.IsUniqueTransaction(stockPosition.Sells, transaction) {
				continue
			}

			if transaction.Type == models.Buytype {
				stockPosition.Quantity += transaction.Quantity
				stockPosition.NetSpend += transaction.Proceeds
				processBuy(ctx, stockPosition, transaction)
			} else {
				stockPosition.Quantity -= transaction.Quantity
				stockPosition.NetSpend -= transaction.Proceeds
				taxMethod := models.FIFO
				sell, err := processSell(ctx, stockPosition, transaction, taxMethod)
				if err != nil {
					return err
				}
				stockPosition.SoldProfit += sell.Profit
				stockPosition.CGTProfit += sell.CGTProfit

				// TODO: Insert tax event
			}
		}

		// Sometimes complete buy and sells dont fully add up to 0
		stockPosition.Quantity = pkgutils.RoundToTwoDecimalPlaces(stockPosition.Quantity)

		if err := s.dbConnector.UpsertStockPosition(ctx, ticker, *stockPosition); err != nil {
			return err
		}
	}

	return nil
}

func processBuy(ctx context.Context, stockPosition *models.StockPosition, transaction models.Transaction) {
	buy := models.Buy{
		Transaction:  transaction,
		QuantityLeft: transaction.Quantity,
	}
	stockPosition.Buys = append(stockPosition.Buys, buy)
}

func processSell(ctx context.Context, stockPosition *models.StockPosition, transaction models.Transaction, taxMethod models.TaxMethod) (*models.Sell, error) {
	// TODO: Use the right algo using taxMethod
	taxProfit := fifo(ctx, transaction, &stockPosition.Buys)

	fmt.Println(taxProfit)

	sell := &models.Sell{
		Transaction: transaction,
		TaxMethod:   taxMethod,
		Profit:      taxProfit.Profit,
		CGTProfit:   taxProfit.CGTProfit,
		BuysSold:    taxProfit.BuysSold,
	}
	stockPosition.Sells = append(stockPosition.Sells, *sell)

	return sell, nil
}
