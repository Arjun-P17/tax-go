// Contains all the repository models

package repository

import (
	"time"
)

/*
	Transaction Models
*/

type TransactionType string

var (
	Buytype  TransactionType = "BUY"
	Selltype TransactionType = "SELL"
)

type TaxMethod string

var (
	FIFO    TaxMethod = "FIFO"
	LIFO    TaxMethod = "LIFO"
	MaxLoss TaxMethod = "MAX_LOSS"
	MinGain TaxMethod = "MIN_GAIN"
	MinCGT  TaxMethod = "MIN_CGT"
)

// Transaction represents the transaction object parsed from the broker.
// It is the base type for Buy and Sell and all numbers are positive.
// Relevant values are adjusted to accomodate stock splits
type Transaction struct {
	ID         string
	Ticker     string
	Currency   string
	Date       time.Time
	Type       TransactionType
	Quantity   float64
	TradePrice float64
	// Adjusted trade price for stock splits
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
	// List of buys (can be partial) that the sell corresponds to
	BuysSold []BuySold
}

func (s Sell) GetDate() time.Time {
	return s.Date
}

func (s Sell) GetBasis() float64 {
	return s.Basis
}

// BuySold represents a buy transaction on each sell
type BuySold struct {
	BuyID    string
	Quantity float64
}

// Buy represents a buy transaction.
type Buy struct {
	Transaction
	QuantityLeft float64
}

func (b Buy) GetDate() time.Time {
	return b.Date
}

func (b Buy) GetBasis() float64 {
	return b.Basis
}

// StockTransactions represents the transactions for a ticker.
type StockTransactions struct {
	Ticker       string
	Transactions []Transaction
}

/*
	Portfolio Models
*/

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

/*
	Tax Models
*/

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
