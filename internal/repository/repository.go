package repository

import (
	"context"
	"errors"

	"github.com/Arjun-P17/tax-go/pkg/configmap"
	"go.mongodb.org/mongo-driver/mongo"
)

type RepositoryInterface interface {
	GetCollection(collection string) *mongo.Collection

	GetAllStockPositions(ctx context.Context) ([]StockPosition, error)
	GetStockPositionOrDefault(ctx context.Context, ticker string) (*StockPosition, error)
	UpsertStockPosition(ctx context.Context, ticker string, stockPosition StockPosition) error

	GetAllStockTransactions(ctx context.Context) ([]StockTransactions, error)
	InsertTransaction(ctx context.Context, transaction Transaction) error

	InsertTaxEvent(ctx context.Context, taxEvent TaxEvent, USDAUD float64) error
}

type Repository struct {
	RepositoryInterface
	config configmap.Database
	client *mongo.Client
}

func NewRepository(client *mongo.Client, dbConfig configmap.Database) (*Repository, error) {
	if client == nil {
		return nil, errors.New("mongodb client is nil")
	}

	return &Repository{
		client: client,
		config: dbConfig,
	}, nil
}

func (c *Repository) GetCollection(collection string) *mongo.Collection {
	client := c.client
	col := client.Database(c.config.DatabaseName).Collection(collection)

	return col
}
