package repository

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

const (
	dbName                 = "tax"
	transactionsCollection = "transactions"
	positionsCollection    = "positions"
)

type Connector struct {
	client *mongo.Client
}

func NewConnector(client *mongo.Client) (*Connector, error) {
	if client == nil {
		return nil, errors.New("mongodb client is nil")
	}

	return &Connector{
		client: client,
	}, nil
}

func (c *Connector) GetCollection(db string, collection string) *mongo.Collection {
	client := c.client
	col := client.Database(db).Collection(collection)

	return col
}
