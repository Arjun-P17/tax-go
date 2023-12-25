package models

import "time"

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

type Sell struct {
	Transaction
	TaxMethod TaxMethod
	Profit    float64
	CGTProfit float64
	BuysSold  []BuySold
}

type BuySold struct {
	BuyID    string
	Quantity float64
}

type Buy struct {
	Transaction
	QuantityLeft float64
}

type StockTransactions struct {
	Ticker       string
	Transactions []Transaction
}

type StockPosition struct {
	Ticker     string
	Quantity   float64
	NetSpend   float64
	SoldProfit float64
	CGTProfit  float64
	Buys       []Buy
	Sells      []Sell
}
