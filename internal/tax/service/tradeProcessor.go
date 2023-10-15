package service

import (
	"context"

	"github.com/Arjun-P17/tax-go/models"
	"github.com/Arjun-P17/tax-go/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (s *Service) processBuy(ctx context.Context, stockPosition *models.StockPosition, transaction models.Transaction) {
	buy := models.Buy{
		Transaction:  transaction,
		QuantityLeft: transaction.Quantity,
	}
	stockPosition.Buys = append(stockPosition.Buys, buy)
}

func (s *Service) processSell(ctx context.Context, stockPosition *models.StockPosition, transaction models.Transaction, taxMethod models.TaxMethod) (*models.Sell, error) {
	// Get these values by doing a tax algo
	profit := 0.0
	cgtProfit := 0.0
	buys := make([]primitive.ObjectID, 0)

	sell := &models.Sell{
		Transaction: transaction,
		TaxMethod:   taxMethod,
		Profit:      profit,
		CGTProfit:   cgtProfit,
		Buys:        buys,
	}
	stockPosition.Sells = append(stockPosition.Sells, *sell)

	return sell, nil
}

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
				stockPosition.NetSpend += transaction.Basis
				s.processBuy(ctx, stockPosition, transaction)
			} else {
				stockPosition.Quantity -= transaction.Quantity
				stockPosition.SoldProfit -= transaction.Basis
				taxMethod := models.FIFO
				sell, err := s.processSell(ctx, stockPosition, transaction, taxMethod)
				if err != nil {
					return err
				}
				stockPosition.SoldProfit += sell.Profit
				stockPosition.CGTProfit += sell.CGTProfit

				// Insert tax event
			}
		}

		// Sometimes complete buy and sells dont fully add up to 0
		stockPosition.Quantity = utils.RoundToTwoDecimalPlaces(stockPosition.Quantity)

		if err := s.dbConnector.UpsertStockPosition(ctx, ticker, *stockPosition); err != nil {
			return err
		}
	}

	return nil
}
