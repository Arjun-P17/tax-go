package mongodb

import (
	"errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type Connetor struct {
	client *mongo.Client
}

func NewConnector(client *mongo.Client) (*Connetor, error) {
	if client == nil {
		return nil, errors.New("mongodb client is nil")
	}

	return &Connetor{
		client: client,
	}, nil
}
