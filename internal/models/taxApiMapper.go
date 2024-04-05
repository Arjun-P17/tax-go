package models

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Arjun-P17/tax-go/proto/go/stockpb"
)

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
		FinancialYear:   st.FinancialYear,
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
