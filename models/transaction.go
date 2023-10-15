package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	ID           primitive.ObjectID
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
	// List of buyIDs that the sell corresponds to
	Buys []primitive.ObjectID
}

// Buy represents a buy transaction.
type Buy struct {
	Transaction
	QuantityLeft float64
}

// StockTransactions represents the transactions for a ticker.
type StockTransactions struct {
	Ticker       string
	Transactions []Transaction
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

func (b Buy) GetDate() time.Time {
	return b.Date
}

func (b Buy) GetBasis() float64 {
	return b.Basis
}

func (s Sell) GetDate() time.Time {
	return s.Date
}

func (s Sell) GetBasis() float64 {
	return s.Basis
}
