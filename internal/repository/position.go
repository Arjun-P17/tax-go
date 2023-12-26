package repository

import (
	"context"
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Connector) GetAllStockPositions(ctx context.Context) ([]StockPosition, error) {
	collection := c.GetCollection(dbName, positionsCollection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stockPositions []StockPosition
	if err = cursor.All(ctx, &stockPositions); err != nil {
		return nil, err
	}

	return stockPositions, nil
}

func (c *Connector) GetStockPositionOrDefault(ctx context.Context, ticker string) (*StockPosition, error) {
	collection := c.GetCollection(dbName, positionsCollection)

	filter := bson.M{"_id": ticker}
	stockPosition := &StockPosition{}
	err := collection.FindOne(ctx, filter).Decode(stockPosition)
	if err == mongo.ErrNoDocuments {
		return &StockPosition{Ticker: ticker}, nil
	} else if err != nil {
		return nil, err
	}

	return stockPosition, nil
}

func (c *Connector) UpsertStockPosition(ctx context.Context, ticker string, stockPosition StockPosition) error {
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
