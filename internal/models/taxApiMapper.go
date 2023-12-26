package models

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Arjun-P17/tax-go/proto/go/stockpb"
)

/*
	Transaction Mappers
*/

func (tt TransactionType) ToProtoModel() stockpb.TransactionType {
	switch tt {
	case Buytype:
		return stockpb.TransactionType_BUY
	case Selltype:
		return stockpb.TransactionType_SELL
	default:
		return stockpb.TransactionType_BUY
	}
}

func (tm TaxMethod) ToProtoModel() stockpb.TaxMethod {
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

func (t Transaction) ToProtoModel() stockpb.Transaction {
	return stockpb.Transaction{
		Id:           t.ID,
		Ticker:       t.Ticker,
		Currency:     t.Currency,
		Date:         &timestamppb.Timestamp{Seconds: t.Date.Unix()},
		Type:         t.Type.ToProtoModel(),
		Quantity:     t.Quantity,
		TradePrice:   t.TradePrice,
		RealPrice:    t.RealPrice,
		Proceeds:     t.Proceeds,
		BrokerageFee: t.BrokerageFee,
		Basis:        t.Basis,
		BrokerProfit: t.BrokerProfit,
		UsdAud:       t.USDAUD,
		SplitFactor:  t.Splitfactor,
	}
}

func (bs BuySold) ToProtoModel() stockpb.BuySold {
	return stockpb.BuySold{
		BuyId:    bs.BuyID,
		Quantity: bs.Quantity,
	}
}

func (s Sell) ToProtoModel() stockpb.Sell {
	var buysSoldProto []*stockpb.BuySold
	for _, repoBuySold := range s.BuysSold {
		bs := repoBuySold.ToProtoModel()
		buysSoldProto = append(buysSoldProto, &bs)
	}

	t := s.Transaction.ToProtoModel()
	return stockpb.Sell{
		Transaction: &t,
		TaxMethod:   s.TaxMethod.ToProtoModel(),
		Profit:      s.Profit,
		CgtProfit:   s.CGTProfit,
		BuysSold:    buysSoldProto,
	}
}

func (b Buy) ToProtoModel() stockpb.Buy {
	t := b.Transaction.ToProtoModel()
	return stockpb.Buy{
		Transaction:  &t,
		QuantityLeft: b.QuantityLeft,
	}
}

/*
	Portfolio Mappers (from Model to Proto)
*/

func (sp StockPosition) ToProtoModel() stockpb.StockPosition {
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
		CgtProfit:  sp.CGTProfit,
		Buys:       buys,
		Sells:      sells,
	}
}

/*
	Tax Mappers (from Model to Proto)
*/

func (te TaxEvent) ToProtoModel() stockpb.TaxEvent {
	return stockpb.TaxEvent{
		Date:         &timestamppb.Timestamp{Seconds: te.Date.Unix()},
		Ticker:       te.Ticker,
		Profit:       te.Profit,
		ProfitAud:    te.ProfitAUD,
		CgtProfit:    te.CGTProfit,
		CgtProfitAud: te.CGTProfitAUD,
	}
}

func (st StockTax) ToProtoModel() stockpb.StockTax {
	var events []*stockpb.TaxEvent
	for _, e := range st.Events {
		protoEvent := e.ToProtoModel()
		events = append(events, &protoEvent)
	}

	return stockpb.StockTax{
		Ticker:          st.Ticker,
		NetProfitCgt:    st.NetProfitCGT,
		NetProfitCgtAud: st.NetProfitCGTAUD,
		NetProfit:       st.NetProfit,
		NetProfitAud:    st.NetProfitAUD,
		GainsCgt:        st.GainsCGT,
		GainsCgtAud:     st.GainsCGTAUD,
		Gains:           st.Gains,
		Losses:          st.Losses,
		Events:          events,
	}
}
