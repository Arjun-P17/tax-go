.PHONY: trades
trades:
	go run internal/trades/main.go

.PHONY: tax
tax:
	go run internal/tax/main.go