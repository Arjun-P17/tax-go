package models

import "time"

// TaxEvent represents a tax event.
type TaxEvent struct {
	Date         time.Time
	Ticker       string
	Profit       float64
	ProfitAUD    float64
	CGTProfit    float64
	CGTProfitAUD float64
}

// StockTax represents stock tax information.
type StockTax struct {
	Ticker          string
	NetProfitCGT    float64
	NetProfitCGTAUD float64
	NetProfit       float64
	NetProfitAUD    float64
	GainsCGT        float64
	GainsCGTAUD     float64
	Gains           float64
	Losses          float64
	Events          []TaxEvent
}
