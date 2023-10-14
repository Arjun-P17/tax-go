package repository

import (
	"context"

	"github.com/Arjun-P17/tax-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Connector) upsertStockTransaction(ctx context.Context, filter bson.M, update bson.M) error {
	collection := c.GetCollection("tax", "transactions")

	options := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, update, options)
	return err
}

func (c *Connector) InsertTransaction(ctx context.Context, transaction models.Transaction) error {
	filter := bson.M{"ticker": transaction.Ticker}

	// Check if the document exists.
	count, err := c.client.Database("tax").Collection("transactions").CountDocuments(ctx, filter)
	if err != nil {
		return err
	}

	if count == 0 {
		newStockTransaction := models.StockTransactions{
			Ticker:       transaction.Ticker,
			Transactions: []models.Transaction{transaction},
		}

		// Insert the new document.
		return c.upsertStockTransaction(ctx, filter, bson.M{"$set": newStockTransaction})
	}

	// If the document exists, push the new transaction into the array.
	update := bson.M{
		"$push": bson.M{"transactions": transaction},
	}

	return c.upsertStockTransaction(ctx, filter, update)
}
