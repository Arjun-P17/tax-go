.PHONY: parseTrades
parseTrades:
	go run internal/apps/parseTrades/main.go

.PHONY: processTrades
processTrades:
	go run internal/apps/processTrades/main.go