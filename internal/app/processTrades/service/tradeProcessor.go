package service

import (
	"context"
	"fmt"

	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/internal/utils"
	pkgutils "github.com/Arjun-P17/tax-go/pkg/utils"
)

func (s *Service) ProcessTrades(ctx context.Context) error {
	allStocksTransactions, err := s.repository.GetAllStockTransactions(ctx)
	if err != nil {
		return err
	}

	for _, stockTransactions := range allStocksTransactions {
		ticker := stockTransactions.Ticker

		stockPosition, err := s.repository.GetStockPositionOrDefault(ctx, ticker)
		if err != nil {
			return err
		}

		for _, transaction := range stockTransactions.Transactions {
			if !utils.IsUniqueTransaction(stockPosition.Buys, transaction) && !utils.IsUniqueTransaction(stockPosition.Sells, transaction) {
				continue
			}

			if transaction.Type == repository.Buytype {
				stockPosition.Quantity += transaction.Quantity
				stockPosition.NetSpend += transaction.Proceeds
				processBuy(ctx, stockPosition, transaction)
			} else {
				stockPosition.Quantity -= transaction.Quantity
				stockPosition.NetSpend -= transaction.Proceeds
				taxMethod := repository.FIFO
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

		if err := s.repository.UpsertStockPosition(ctx, ticker, *stockPosition); err != nil {
			return err
		}
	}

	return nil
}

func processBuy(ctx context.Context, stockPosition *repository.StockPosition, transaction repository.Transaction) {
	buy := repository.Buy{
		Transaction:  transaction,
		QuantityLeft: transaction.Quantity,
	}
	stockPosition.Buys = append(stockPosition.Buys, buy)
}

func processSell(ctx context.Context, stockPosition *repository.StockPosition, transaction repository.Transaction, taxMethod repository.TaxMethod) (*repository.Sell, error) {
	// TODO: Use the right algo using taxMethod
	taxProfit := fifo(ctx, transaction, &stockPosition.Buys)

	fmt.Println(taxProfit)

	sell := &repository.Sell{
		Transaction: transaction,
		TaxMethod:   taxMethod,
		Profit:      taxProfit.Profit,
		CGTProfit:   taxProfit.CGTProfit,
		BuysSold:    taxProfit.BuysSold,
	}
	stockPosition.Sells = append(stockPosition.Sells, *sell)

	return sell, nil
}
