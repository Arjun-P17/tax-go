package models

import "time"

// TaxableEvent represents a single taxable transaction that needs to be reported.
// It contains the profit/loss information and its tax implications.
type TaxableEvent struct {
	Date             time.Time // Date when the taxable event occurred
	Ticker           string    // Stock symbol involved in the event
	Profit           float64   // Total profit/loss in original currency
	ProfitAUD        float64   // Total profit/loss converted to AUD
	TaxableProfit    float64   // Amount subject to tax in original currency
	TaxableProfitAUD float64   // Amount subject to tax in AUD
}

// TaxYearSummary represents the tax summary for a specific financial year.
// It aggregates all taxable events and provides totals for tax reporting purposes.
type TaxYearSummary struct {
	FinancialYear       string         // The financial year (e.g., "2023-24")
	NetTaxableProfit    float64        // Total taxable profit for the year
	NetTaxableProfitAUD float64        // Total taxable profit in AUD
	NetProfit           float64        // Total profit (including non-taxable) for the year
	NetProfitAUD        float64        // Total profit in AUD
	TaxableGains        float64        // Total capital gains subject to tax
	TaxableGainsAUD     float64        // Total capital gains in AUD
	Gains               float64        // Total capital gains (including non-taxable)
	Losses              float64        // Total capital losses
	Events              []TaxableEvent // List of all taxable events for the year
}
