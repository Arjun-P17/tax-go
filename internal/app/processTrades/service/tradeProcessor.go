package service

import (
	"context"

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
			unique := processTransaction(stockPosition, transaction)
			// No need for the unique check but keeping it for explicitness
			if !unique {
				continue
			}

			// TODO: Insert tax event on sells
		}

		// Sometimes complete buy and sells dont fully add up to 0
		stockPosition.Quantity = pkgutils.RoundToTwoDecimalPlaces(stockPosition.Quantity)

		if err := s.repository.UpsertStockPosition(ctx, ticker, *stockPosition); err != nil {
			return err
		}
	}

	return nil
}

// processTransaction processes the transaction and updates the stock position
// returns false if the transaction is not unique
func processTransaction(stockPosition *repository.StockPosition, transaction repository.Transaction) bool {
	if !utils.IsUniqueTransaction(stockPosition.Buys, transaction) && !utils.IsUniqueTransaction(stockPosition.Sells, transaction) {
		return false
	}

	if transaction.Type == repository.Buytype {
		processBuy(stockPosition, transaction)
	} else {
		taxMethod := repository.FIFO
		sell := processSell(stockPosition, transaction, taxMethod)

		stockPosition.SoldProfit += sell.Profit
		stockPosition.CGTProfit += sell.CGTProfit
	}
	return true
}

func processBuy(stockPosition *repository.StockPosition, transaction repository.Transaction) {
	stockPosition.Quantity += transaction.Quantity
	stockPosition.NetSpend += transaction.Proceeds
	buy := repository.Buy{
		Transaction:  transaction,
		QuantityLeft: transaction.Quantity,
	}
	stockPosition.Buys = append(stockPosition.Buys, buy)
}

func processSell(stockPosition *repository.StockPosition, transaction repository.Transaction, taxMethod repository.TaxMethod) *repository.Sell {
	stockPosition.Quantity -= transaction.Quantity
	stockPosition.NetSpend -= transaction.Proceeds

	// TODO: Use the right algo using taxMethod
	taxProfit := fifo(transaction, &stockPosition.Buys)

	sell := &repository.Sell{
		Transaction: transaction,
		TaxMethod:   taxMethod,
		Profit:      taxProfit.Profit,
		CGTProfit:   taxProfit.CGTProfit,
		BuysSold:    taxProfit.BuysSold,
	}
	stockPosition.Sells = append(stockPosition.Sells, *sell)

	return sell
}
