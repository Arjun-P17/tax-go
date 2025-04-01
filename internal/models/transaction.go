package models

import "time"

// TradeType represents the type of trade transaction (buy or sell).
// This is used to distinguish between purchase and sale transactions
// for tax calculation purposes.
type TradeType string

var (
	Buytype  TradeType = "BUY"  // Represents a purchase transaction
	Selltype TradeType = "SELL" // Represents a sale transaction
)

// TaxCalculationMethod defines the strategy used to calculate capital gains tax.
// Different methods can be used to match buy and sell transactions for tax purposes.
type TaxCalculationMethod string

var (
	FIFO    TaxCalculationMethod = "FIFO"     // First In, First Out - oldest shares are sold first
	LIFO    TaxCalculationMethod = "LIFO"     // Last In, First Out - newest shares are sold first
	MaxLoss TaxCalculationMethod = "MAX_LOSS" // Maximizes capital losses for tax purposes
	MinGain TaxCalculationMethod = "MIN_GAIN" // Minimizes capital gains for tax purposes
	MinCGT  TaxCalculationMethod = "MIN_CGT"  // Minimizes capital gains tax liability
)

// StockTransactions groups all transactions for a specific stock ticker.
// This structure is used to organize and process transactions by stock symbol.
type StockTransactions struct {
	Ticker       string             // The stock symbol (e.g., AAPL, GOOGL)
	Transactions []TradeTransaction // List of all transactions for this stock
}

// TradeTransaction represents a single trade execution in the system.
// It contains all the necessary information for tax calculations and portfolio tracking.
type TradeTransaction struct {
	ID             string    // Unique identifier for the transaction
	Ticker         string    // Stock symbol being traded
	Currency       string    // Currency of the transaction (e.g., USD, AUD)
	Date           time.Time // When the transaction occurred
	Type           TradeType // Whether this is a buy or sell transaction
	Quantity       float64   // Number of shares traded
	ExecutionPrice float64   // Price per share at execution
	RealPrice      float64   // Actual price after adjustments (e.g., splits)
	Proceeds       float64   // Total value of the transaction (Quantity * ExecutionPrice)
	Commission     float64   // Brokerage fees for the transaction
	Basis          float64   // Cost basis for tax calculations
	BrokerProfit   float64   // Profit/loss from the broker's perspective
	USDAUD         float64   // Exchange rate between USD and AUD at transaction time
	Splitfactor    float64   // Stock split factor (1.0 for no split)
}

// SellTransaction extends TradeTransaction with tax-specific information for sales.
// It tracks how the sale was matched with previous purchases for tax purposes.
type SellTransaction struct {
	TradeTransaction
	TaxCalcMethod    TaxCalculationMethod // Method used to calculate capital gains
	Profit           float64              // Total profit from the sale
	TaxableProfit    float64              // Profit subject to capital gains tax
	MatchedPurchases []MatchedPurchase    // List of buy transactions matched to this sale
}

// MatchedPurchase represents a link between a sell transaction and its corresponding buy transaction.
// This is used for tax calculation purposes to determine capital gains/losses.
type MatchedPurchase struct {
	BuyID    string  // ID of the matched buy transaction
	Quantity float64 // Quantity of shares matched from this buy
}

// BuyTransaction extends TradeTransaction with tracking of remaining shares.
// This is used to track how many shares from a purchase are still available for matching with future sales.
type BuyTransaction struct {
	TradeTransaction
	QuantityLeft float64 // Number of shares from this buy that haven't been sold
}

// PortfolioPosition represents the current state of a stock position in a portfolio.
// It tracks all buy and sell transactions for a specific stock and calculates current position metrics.
type PortfolioPosition struct {
	Ticker        string            // Stock symbol
	Quantity      float64           // Current number of shares held
	NetSpend      float64           // Total amount spent on current position
	SoldProfit    float64           // Total profit from closed positions
	TaxableProfit float64           // Total taxable profit from closed positions
	Buys          []BuyTransaction  // List of all buy transactions
	Sells         []SellTransaction // List of all sell transactions
}
