package models

import (
	"time"
)

// StockTransactions represents stock transactions.
type StockTransactions struct {
	Ticker       string
	Transactions []Transaction
}

type TransactionType string

var (
	Buytype  TransactionType = "Buy"
	Selltype TransactionType = "Sell"
)

type TaxMethod string

var (
	FIFO    TaxMethod = "FIFO"
	LIFO    TaxMethod = "LIFO"
	MaxLoss TaxMethod = "MaxLoss"
	MinGain TaxMethod = "MinGain"
	MinCGT  TaxMethod = "MinCGT"
)

// Transaction represents the transaction object parsed from the broker.
// It is the base type for Buy and Sell and all numbers are positive
type Transaction struct {
	ID           string
	Ticker       string
	Currency     string
	Date         time.Time
	Type         TransactionType
	Quantity     float64
	TradePrice   float64
	RealPrice    float64
	Proceeds     float64
	BrokerageFee float64
	Basis        float64
	BrokerProfit float64
	USDAUD       float64
	Splitfactor  float64
}

// Sell represents a sell transaction.
type Sell struct {
	Transaction
	TaxMethod TaxMethod
	Profit    float64
	CGTProfit float64
	BuyEvents []BuyEvent
}

// BuyEvent represents a buy that all or part of the sell corresponds to
type BuyEvent struct {
	BuyID     string
	Quantity  float64
	RealPrice float64
	Date      time.Time
}

// Buy represents a buy transaction.
type Buy struct {
	Transaction
	QuantityLeft float64
	SellEvents   []SellEvent
}

// SellEvent represents a sell that all or part of the buy corresponds to
type SellEvent struct {
	SellID    string
	Quantity  float64
	RealPrice float64
	Date      time.Time
}

// StockPosition represents a portfolio position.
type StockPosition struct {
	Ticker     string
	Quantity   float64
	NetSpend   float64
	SoldProfit float64
	CGTProfit  float64
	Buys       []Buy
	Sells      []Sell
}
