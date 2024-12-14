package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"sort"

	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/pkg/configmap"
	"github.com/Arjun-P17/tax-go/pkg/mongodb"
)

var ticker = flag.String("ticker", "", "the ticker to get active trades for")

func main() {
	flag.Parse()
	if ticker == nil {
		log.Fatal("ticker is nil")
	}
	fmt.Println("Displaying active buys for ticker:", *ticker)

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

	repo, err := repository.NewRepository(client, config.Database)
	if err != nil {
		log.Fatal(err)
	}

	position, err := repo.GetStockPositionOrDefault(ctx, *ticker)
	if err != nil {
		log.Fatal(err)
	}

	// Display the trades
	buys := position.Buys
	sort.Slice(buys, func(i, j int) bool {
		return buys[i].Date.Before(buys[j].Date)
	})

	for _, buy := range buys {
		if buy.QuantityLeft > 0 {
			displayFmt := "Ticker %s, Buy Date %s, Buy Price %.2f, Quantity %.2f, Quantity Left %.2f"
			fmt.Println(fmt.Sprintf(displayFmt, buy.Ticker, buy.Date, buy.RealPrice, buy.Quantity, buy.QuantityLeft))
		}
	}
}
