package models

import "time"

type TaxableEvent struct {
	Date             time.Time
	Ticker           string
	Profit           float64
	ProfitAUD        float64
	TaxableProfit    float64
	TaxableProfitAUD float64
}

type TaxYearSummary struct {
	FinancialYear       string
	NetTaxableProfit    float64
	NetTaxableProfitAUD float64
	NetProfit           float64
	NetProfitAUD        float64
	TaxableGains        float64
	TaxableGainsAUD     float64
	Gains               float64
	Losses              float64
	Events              []TaxableEvent
}
