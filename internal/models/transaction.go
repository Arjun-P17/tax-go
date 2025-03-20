package models

import "time"

type TradeType string

var (
	Buytype  TradeType = "BUY"
	Selltype TradeType = "SELL"
)

type TaxCalculationMethod string

var (
	FIFO    TaxCalculationMethod = "FIFO"
	LIFO    TaxCalculationMethod = "LIFO"
	MaxLoss TaxCalculationMethod = "MAX_LOSS"
	MinGain TaxCalculationMethod = "MIN_GAIN"
	MinCGT  TaxCalculationMethod = "MIN_CGT"
)

type StockTransactions struct {
	Ticker       string
	Transactions []TradeTransaction
}

type TradeTransaction struct {
	ID             string
	Ticker         string
	Currency       string
	Date           time.Time
	Type           TradeType
	Quantity       float64
	ExecutionPrice float64
	RealPrice      float64
	Proceeds       float64
	Commission     float64
	Basis          float64
	BrokerProfit   float64
	USDAUD         float64
	Splitfactor    float64
}

type SellTransaction struct {
	TradeTransaction
	TaxCalcMethod    TaxCalculationMethod
	Profit           float64
	TaxableProfit    float64
	MatchedPurchases []MatchedPurchase
}

type MatchedPurchase struct {
	BuyID    string
	Quantity float64
}

type BuyTransaction struct {
	TradeTransaction
	QuantityLeft float64
}

type PortfolioPosition struct {
	Ticker        string
	Quantity      float64
	NetSpend      float64
	SoldProfit    float64
	TaxableProfit float64
	Buys          []BuyTransaction
	Sells         []SellTransaction
}
