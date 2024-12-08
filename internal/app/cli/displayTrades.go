package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"slices"

	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/pkg/configmap"
	"github.com/Arjun-P17/tax-go/pkg/mongodb"
)

var ticker = flag.String("ticker", "", "the ticker to get active trades for")

func main() {
	flag.Parse()

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

	if ticker == nil {
		log.Fatal("ticker is nil")
	}

	position, err := repo.GetStockPositionOrDefault(ctx, *ticker)
	if err != nil {
		log.Fatal(err)
	}

	slices.


	buys := position.Buys
	for _, buy := range buys {

	}
}
