package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/Arjun-P17/tax-go/internal/models"
)

func (c *Connector) GetAllStockTransactions(ctx context.Context) ([]models.StockTransactions, error) {
	collection := c.GetCollection(dbName, transactionsCollection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stockTransactions []models.StockTransactions
	if err = cursor.All(ctx, &stockTransactions); err != nil {
		return nil, err
	}

	return stockTransactions, nil
}

func (c *Connector) upsertStockTransaction(ctx context.Context, filter bson.M, update bson.M) error {
	collection := c.GetCollection(dbName, transactionsCollection)

	options := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, update, options)
	return err
}

func (c *Connector) InsertTransaction(ctx context.Context, transaction models.Transaction) error {
	collection := c.GetCollection(dbName, transactionsCollection)

	// Check if the document exists.
	filter := bson.M{"ticker": transaction.Ticker}
	count, err := collection.CountDocuments(ctx, filter)
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
		"$addToSet": bson.M{"transactions": transaction},
	}

	return c.upsertStockTransaction(ctx, filter, update)
}
