package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Arjun-P17/tax-go/internal/trades/service.go"
	"github.com/Arjun-P17/tax-go/pkg/configmap"
	"github.com/Arjun-P17/tax-go/pkg/mongodb"
)

const configPath = "config.yaml"

func main() {
	config, err := configmap.ReadConfigFile(configPath)
	if err != nil {
		log.Fatal(err)
	}

	mongoURI := fmt.Sprintf("mongodb://%s:%d", config.Database.Host, config.Database.Port)
	opts := mongodb.Options{
		URI: mongoURI,
	}
	client, err := mongodb.NewClient(context.Background(), &opts)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	transactions, err := service.ParseTransactions(config.Trades.CSVPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, transaction := range transactions {
		fmt.Println(transaction.Ticker)
	}
}
