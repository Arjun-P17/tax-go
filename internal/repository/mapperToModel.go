// Mappers from repository to internal models

package repository

import (
	"github.com/Arjun-P17/tax-go/internal/models"
)

/*
	Transaction Mappers
*/

// Returns BuyType as default if not found
func (tt TransactionType) ToModel() models.TradeType {
	switch tt {
	case Buytype:
		return models.Buytype
	case Selltype:
		return models.Selltype
	default:
		return models.Buytype
	}
}

// Returns FIFO as default if not found
func (tm TaxMethod) ToModel() models.TaxCalculationMethod {
	switch tm {
	case FIFO:
		return models.FIFO
	case LIFO:
		return models.LIFO
	case MaxLoss:
		return models.MaxLoss
	case MinGain:
		return models.MinGain
	case MinCGT:
		return models.MinCGT
	default:
		return models.FIFO
	}
}

func (t Transaction) ToModel() models.TradeTransaction {
	return models.TradeTransaction{
		ID:             t.ID,
		Ticker:         t.Ticker,
		Currency:       t.Currency,
		Date:           t.Date,
		Type:           t.Type.ToModel(),
		Quantity:       t.Quantity,
		ExecutionPrice: t.TradePrice,
		RealPrice:      t.RealPrice,
		Proceeds:       t.Proceeds,
		Commission:     t.BrokerageFee,
		Basis:          t.Basis,
		BrokerProfit:   t.BrokerProfit,
		USDAUD:         t.USDAUD,
		Splitfactor:    t.Splitfactor,
	}
}

func (bs BuySold) ToModel() models.MatchedPurchase {
	return models.MatchedPurchase{
		BuyID:    bs.BuyID,
		Quantity: bs.Quantity,
	}
}

func (s Sell) ToModel() models.SellTransaction {
	var buysSoldModel []models.MatchedPurchase
	for _, bs := range s.BuysSold {
		buysSoldModel = append(buysSoldModel, bs.ToModel())
	}

	return models.SellTransaction{
		TradeTransaction: s.Transaction.ToModel(),
		TaxCalcMethod:    s.TaxMethod.ToModel(),
		Profit:           s.Profit,
		TaxableProfit:    s.CGTProfit,
		MatchedPurchases: buysSoldModel,
	}
}

func (b Buy) ToModel() models.BuyTransaction {
	return models.BuyTransaction{
		TradeTransaction: b.Transaction.ToModel(),
		QuantityLeft:     b.QuantityLeft,
	}
}

/*
	Portfolio Mappers
*/

func (sp StockPosition) ToModel() models.PortfolioPosition {
	var buysModel []models.BuyTransaction
	for _, buy := range sp.Buys {
		buysModel = append(buysModel, buy.ToModel())
	}

	var sellsModel []models.SellTransaction
	for _, sell := range sp.Sells {
		sellsModel = append(sellsModel, sell.ToModel())
	}

	return models.PortfolioPosition{
		Ticker:        sp.Ticker,
		Quantity:      sp.Quantity,
		NetSpend:      sp.NetSpend,
		SoldProfit:    sp.SoldProfit,
		TaxableProfit: sp.CGTProfit,
		Buys:          buysModel,
		Sells:         sellsModel,
	}
}

/*
	Tax Mappers
*/

func (te TaxEvent) ToModel() models.TaxableEvent {
	return models.TaxableEvent{
		Date:             te.Date,
		Ticker:           te.Ticker,
		Profit:           te.Profit,
		ProfitAUD:        te.ProfitAUD,
		TaxableProfit:    te.CGTProfit,
		TaxableProfitAUD: te.CGTProfitAUD,
	}
}

func (st StockTax) ToModel() models.TaxYearSummary {
	var eventsModel []models.TaxableEvent
	for _, event := range st.Events {
		eventsModel = append(eventsModel, event.ToModel())
	}

	return models.TaxYearSummary{
		FinancialYear:       st.FinancialYear,
		NetTaxableProfit:    st.NetProfitCGT,
		NetTaxableProfitAUD: st.NetProfitCGTAUD,
		NetProfit:           st.NetProfit,
		NetProfitAUD:        st.NetProfitAUD,
		TaxableGains:        st.GainsCGT,
		TaxableGainsAUD:     st.GainsCGTAUD,
		Gains:               st.Gains,
		Losses:              st.Losses,
		Events:              eventsModel,
	}
}
