package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/Arjun-P17/tax-go/internal/api"
	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/internal/service"
	"github.com/Arjun-P17/tax-go/pkg/configmap"
	"github.com/Arjun-P17/tax-go/pkg/mongodb"
	"github.com/Arjun-P17/tax-go/proto/go/stockpb"
	"google.golang.org/grpc"
)

// TODO: make cleaner
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

	db, err := repository.NewConnector(client)
	if err != nil {
		log.Fatal(err)
	}

	service, err := service.NewService(db)

	api, err := api.NewApi(service)

	// Create and start grpc server
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var grpcOpts []grpc.ServerOption
	grpcServer := grpc.NewServer(grpcOpts...)

	stockpb.RegisterStockServiceServer(grpcServer, api)

	grpcServer.Serve(lis)

}
