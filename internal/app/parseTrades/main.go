package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Arjun-P17/tax-go/internal/app/parseTrades/service"
	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/pkg/configmap"
	"github.com/Arjun-P17/tax-go/pkg/mongodb"
)

func main() {
	ctx := context.Background()

	config, err := configmap.ReadConfigFile(configmap.ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%d", config.Database.Host, config.Database.Port)
	opts := mongodb.Options{
		URI: mongoURI,
	}
	client, err := mongodb.NewClient(ctx, &opts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(ctx); err != nil {
			panic(err)
		}
	}()

	repo, err := repository.NewRepository(client)
	if err != nil {
		log.Fatal(err)
	}

	transactions, err := service.ParseTransactions(config.Trades.CSVPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, transaction := range transactions {
		fmt.Println(transaction.Ticker)
		repo.InsertTransaction(ctx, *transaction)
	}
}
