.PHONY: parseTrades
parseTrades:
	go run internal/app/parseTrades/main.go

.PHONY: processTrades
processTrades:
	go run internal/app/processTrades/main.go