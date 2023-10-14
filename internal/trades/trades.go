package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Arjun-P17/tax-go/internal/trades/service"
	"github.com/Arjun-P17/tax-go/pkg/configmap"
	"github.com/Arjun-P17/tax-go/pkg/mongodb"
	"github.com/Arjun-P17/tax-go/repository"
)

const configPath = "config.yaml"

func main() {
	ctx := context.Background()

	config, err := configmap.ReadConfigFile(configPath)
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

	connector, err := repository.NewConnector(client)
	if err != nil {
		log.Fatal(err)
	}

	transactions, err := service.ParseTransactions(config.Trades.CSVPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, transaction := range transactions {
		fmt.Println(transaction.Ticker)
		connector.InsertTransaction(ctx, *transaction)
	}
}
