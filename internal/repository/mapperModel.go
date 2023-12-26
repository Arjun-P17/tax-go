// Mappers from repository to internal models

package repository

import (
	"github.com/Arjun-P17/tax-go/internal/models"
)

/*
	Transaction Mappers
*/

// Returns BuyType as default if not found
func (tt TransactionType) ToModel() models.TransactionType {
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
func (tm TaxMethod) ToModel() models.TaxMethod {
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

func (t Transaction) ToModel() models.Transaction {
	return models.Transaction{
		ID:           t.ID,
		Ticker:       t.Ticker,
		Currency:     t.Currency,
		Date:         t.Date,
		Type:         t.Type.ToModel(),
		Quantity:     t.Quantity,
		TradePrice:   t.TradePrice,
		RealPrice:    t.RealPrice,
		Proceeds:     t.Proceeds,
		BrokerageFee: t.BrokerageFee,
		Basis:        t.Basis,
		BrokerProfit: t.BrokerProfit,
		USDAUD:       t.USDAUD,
		Splitfactor:  t.Splitfactor,
	}
}

func (bs BuySold) ToModel() models.BuySold {
	return models.BuySold{
		BuyID:    bs.BuyID,
		Quantity: bs.Quantity,
	}
}

func (s Sell) ToModel() models.Sell {
	var buysSoldModel []models.BuySold
	for _, bs := range s.BuysSold {
		buysSoldModel = append(buysSoldModel, bs.ToModel())
	}

	return models.Sell{
		Transaction: s.Transaction.ToModel(),
		TaxMethod:   s.TaxMethod.ToModel(),
		Profit:      s.Profit,
		CGTProfit:   s.CGTProfit,
		BuysSold:    buysSoldModel,
	}
}

func (b Buy) ToModel() models.Buy {
	return models.Buy{
		Transaction:  b.Transaction.ToModel(),
		QuantityLeft: b.QuantityLeft,
	}
}

/*
	Portfolio Mappers
*/

func (sp StockPosition) ToModel() models.StockPosition {
	var buysModel []models.Buy
	for _, buy := range sp.Buys {
		buysModel = append(buysModel, buy.ToModel())
	}

	var sellsModel []models.Sell
	for _, sell := range sp.Sells {
		sellsModel = append(sellsModel, sell.ToModel())
	}

	return models.StockPosition{
		Ticker:     sp.Ticker,
		Quantity:   sp.Quantity,
		NetSpend:   sp.NetSpend,
		SoldProfit: sp.SoldProfit,
		CGTProfit:  sp.CGTProfit,
		Buys:       buysModel,
		Sells:      sellsModel,
	}
}

/*
	Tax Mappers
*/

func (te TaxEvent) ToModel() models.TaxEvent {
	return models.TaxEvent{
		Date:         te.Date,
		Ticker:       te.Ticker,
		Profit:       te.Profit,
		ProfitAUD:    te.ProfitAUD,
		CGTProfit:    te.CGTProfit,
		CGTProfitAUD: te.CGTProfitAUD,
	}
}

func (st StockTax) ToModel() models.StockTax {
	var eventsModel []models.TaxEvent
	for _, event := range st.Events {
		eventsModel = append(eventsModel, event.ToModel())
	}

	return models.StockTax{
		Ticker:          st.Ticker,
		NetProfitCGT:    st.NetProfitCGT,
		NetProfitCGTAUD: st.NetProfitCGTAUD,
		NetProfit:       st.NetProfit,
		NetProfitAUD:    st.NetProfitAUD,
		GainsCGT:        st.GainsCGT,
		GainsCGTAUD:     st.GainsCGTAUD,
		Gains:           st.Gains,
		Losses:          st.Losses,
		Events:          eventsModel,
	}
}
