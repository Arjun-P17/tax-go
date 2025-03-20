package models

import (
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/Arjun-P17/tax-go/proto/go/stockpb"
)

/*
	Tax Mappers (from Model to Proto)
*/

func (te TaxableEvent) ToProtoModel() stockpb.TaxEvent {
	return stockpb.TaxEvent{
		Date:         &timestamppb.Timestamp{Seconds: te.Date.Unix()},
		Ticker:       te.Ticker,
		Profit:       te.Profit,
		ProfitAud:    te.ProfitAUD,
		CgtProfit:    te.TaxableProfit,
		CgtProfitAud: te.TaxableProfitAUD,
	}
}

func (st TaxYearSummary) ToProtoModel() stockpb.StockTax {
	var events []*stockpb.TaxEvent
	for _, e := range st.Events {
		protoEvent := e.ToProtoModel()
		events = append(events, &protoEvent)
	}

	return stockpb.StockTax{
		FinancialYear:   st.FinancialYear,
		NetProfitCgt:    st.NetTaxableProfit,
		NetProfitCgtAud: st.NetTaxableProfitAUD,
		NetProfit:       st.NetProfit,
		NetProfitAud:    st.NetProfitAUD,
		GainsCgt:        st.TaxableGains,
		GainsCgtAud:     st.TaxableGainsAUD,
		Gains:           st.Gains,
		Losses:          st.Losses,
		Events:          events,
	}
}
