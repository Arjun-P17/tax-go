package service

import (
	"context"

	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/internal/utils"
	"github.com/Arjun-P17/tax-go/pkg/date"
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
			_ = s.processTransaction(ctx, stockPosition, transaction)
		}

		// Sometimes complete buy and sells dont fully add up to 0 so round it off
		stockPosition.Quantity = pkgutils.RoundToTwoDecimalPlaces(stockPosition.Quantity)

		if err := s.repository.UpsertStockPosition(ctx, ticker, *stockPosition); err != nil {
			return err
		}
	}

	return nil
}

// processTransaction processes the transaction and updates the stock position
// returns false if the transaction is not unique
func (s *Service) processTransaction(ctx context.Context, stockPosition *repository.StockPosition, transaction repository.Transaction) bool {
	if !utils.IsUniqueTransaction(stockPosition.Buys, transaction) && !utils.IsUniqueTransaction(stockPosition.Sells, transaction) {
		return false
	}

	if transaction.Type == repository.Buytype {
		processBuy(stockPosition, transaction)
	} else {
		taxMethod := repository.FIFO
		s.processSell(ctx, stockPosition, transaction, taxMethod)
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

func (s *Service) processSell(ctx context.Context, stockPosition *repository.StockPosition, transaction repository.Transaction, taxMethod repository.TaxMethod) {
	stockPosition.Quantity -= transaction.Quantity
	stockPosition.NetSpend -= transaction.Proceeds

	// TODO: Use the right algo using taxMethod
	taxProfit := fifo(transaction, &stockPosition.Buys)

	stockPosition.SoldProfit += taxProfit.Profit
	stockPosition.CGTProfit += taxProfit.CGTProfit

	sell := repository.Sell{
		Transaction: transaction,
		TaxMethod:   taxMethod,
		Profit:      taxProfit.Profit,
		CGTProfit:   taxProfit.CGTProfit,
		BuysSold:    taxProfit.BuysSold,
	}
	stockPosition.Sells = append(stockPosition.Sells, sell)

	// Try get USDAUD value from currencyPair map otherwise use 1.5
	USDAUD := 1.5
	if price, ok := s.currencyMap[date.NewDate(transaction.Date)]; ok {
		USDAUD = 1 / price
	}

	// Add tax event
	taxEvent := repository.TaxEvent{
		Ticker:       transaction.Ticker,
		Date:         transaction.Date,
		Sell:         sell,
		Profit:       taxProfit.Profit,
		ProfitAUD:    taxProfit.Profit * USDAUD,
		CGTProfit:    taxProfit.CGTProfit,
		CGTProfitAUD: taxProfit.CGTProfit * USDAUD,
	}
	s.repository.InsertTaxEvent(ctx, taxEvent, USDAUD)
}
