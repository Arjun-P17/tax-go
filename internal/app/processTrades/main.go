package main

import (
	"context"
	"fmt"
	"log"

	"github.com/Arjun-P17/tax-go/internal/app/processTrades/service"
	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/pkg/configmap"
	"github.com/Arjun-P17/tax-go/pkg/currency"
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

	repo, err := repository.NewRepository(client, config.Database)
	if err != nil {
		log.Fatal(err)
	}

	currencyMap, err := currency.GetAUDUSDMap()
	if err != nil {
		log.Fatal(err)
	}

	service := service.NewService(repo, currencyMap)
	service.ProcessTrades(ctx)
}
