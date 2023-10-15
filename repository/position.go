package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/Arjun-P17/tax-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Connector) GetStockPositionOrDefault(ctx context.Context, ticker string) (*models.StockPosition, error) {
	collection := c.GetCollection(dbName, positionsCollection)

	filter := bson.M{"_id": ticker}
	stockPosition := &models.StockPosition{}
	err := collection.FindOne(ctx, filter).Decode(stockPosition)
	if err == mongo.ErrNoDocuments {
		return &models.StockPosition{Ticker: ticker}, nil
	} else if err != nil {
		return nil, err
	}

	return stockPosition, nil
}

func (c *Connector) UpsertStockPosition(ctx context.Context, ticker string, stockPosition models.StockPosition) error {
	collection := c.GetCollection(dbName, positionsCollection)

	filter := bson.M{"ticker": ticker}
	update := bson.M{"$set": stockPosition}
	opts := options.Update().SetUpsert(true)
	result, err := collection.UpdateOne(ctx, filter, update, opts)
	if err != nil {
		return err
	}

	if result.ModifiedCount > 0 {
		fmt.Println("Existing position document updated:", ticker)
	} else if result.UpsertedID != nil {
		fmt.Println("New position document created:", ticker)
	} else {
		return errors.New("Failed to upsert position document:" + ticker)
	}

	return nil
}
