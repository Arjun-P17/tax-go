package models

import (
	"github.com/Arjun-P17/tax-go/proto/go/stockpb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

/*
	Transaction Mappers (from model to proto)
*/

func (tt TradeType) ToProtoModel() stockpb.TransactionType {
	switch tt {
	case Buytype:
		return stockpb.TransactionType_BUY
	case Selltype:
		return stockpb.TransactionType_SELL
	default:
		return stockpb.TransactionType_BUY
	}
}

func (tm TaxCalculationMethod) ToProtoModel() stockpb.TaxMethod {
	switch tm {
	case FIFO:
		return stockpb.TaxMethod_FIFO
	case LIFO:
		return stockpb.TaxMethod_LIFO
	case MaxLoss:
		return stockpb.TaxMethod_MAX_LOSS
	case MinGain:
		return stockpb.TaxMethod_MIN_GAIN
	case MinCGT:
		return stockpb.TaxMethod_MIN_CGT
	default:
		return stockpb.TaxMethod_FIFO
	}
}

func (t TradeTransaction) ToProtoModel() stockpb.Transaction {
	return stockpb.Transaction{
		Id:           t.ID,
		Ticker:       t.Ticker,
		Currency:     t.Currency,
		Date:         &timestamppb.Timestamp{Seconds: t.Date.Unix()},
		Type:         t.Type.ToProtoModel(),
		Quantity:     t.Quantity,
		TradePrice:   t.ExecutionPrice,
		RealPrice:    t.RealPrice,
		Proceeds:     t.Proceeds,
		BrokerageFee: t.Commission,
		Basis:        t.Basis,
		BrokerProfit: t.BrokerProfit,
		UsdAud:       t.USDAUD,
		SplitFactor:  t.Splitfactor,
	}
}

func (bs MatchedPurchase) ToProtoModel() stockpb.BuySold {
	return stockpb.BuySold{
		BuyId:    bs.BuyID,
		Quantity: bs.Quantity,
	}
}

func (s SellTransaction) ToProtoModel() stockpb.Sell {
	var buysSoldProto []*stockpb.BuySold
	for _, repoBuySold := range s.MatchedPurchases {
		bs := repoBuySold.ToProtoModel()
		buysSoldProto = append(buysSoldProto, &bs)
	}

	t := s.TradeTransaction.ToProtoModel()
	return stockpb.Sell{
		Transaction: &t,
		TaxMethod:   s.TaxCalcMethod.ToProtoModel(),
		Profit:      s.Profit,
		CgtProfit:   s.TaxableProfit,
		BuysSold:    buysSoldProto,
	}
}

func (b BuyTransaction) ToProtoModel() stockpb.Buy {
	t := b.TradeTransaction.ToProtoModel()
	return stockpb.Buy{
		Transaction:  &t,
		QuantityLeft: b.QuantityLeft,
	}
}

/*
	Portfolio Mappers (from Model to Proto)
*/

func (sp PortfolioPosition) ToProtoModel() stockpb.StockPosition {
	var buys []*stockpb.Buy
	for _, repoBuy := range sp.Buys {
		t := repoBuy.ToProtoModel()
		buys = append(buys, &t)
	}

	var sells []*stockpb.Sell
	for _, repoSell := range sp.Sells {
		t := repoSell.ToProtoModel()
		sells = append(sells, &t)
	}

	return stockpb.StockPosition{
		Ticker:     sp.Ticker,
		Quantity:   sp.Quantity,
		NetSpend:   sp.NetSpend,
		SoldProfit: sp.SoldProfit,
		CgtProfit:  sp.TaxableProfit,
		Buys:       buys,
		Sells:      sells,
	}
}
