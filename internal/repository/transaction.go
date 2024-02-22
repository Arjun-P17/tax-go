package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Repository) GetAllStockTransactions(ctx context.Context) ([]StockTransactions, error) {
	collection := c.GetCollection(c.config.DatabaseName, c.config.TransactionsCollection)

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var stockTransactions []StockTransactions
	if err = cursor.All(ctx, &stockTransactions); err != nil {
		return nil, err
	}

	return stockTransactions, nil
}

func (c *Repository) upsertStockTransaction(ctx context.Context, filter bson.M, update bson.M) error {
	collection := c.GetCollection(c.config.DatabaseName, c.config.TransactionsCollection)

	options := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, update, options)
	return err
}

func (c *Repository) InsertTransaction(ctx context.Context, transaction Transaction) error {
	collection := c.GetCollection(c.config.DatabaseName, c.config.TransactionsCollection)

	// Check if the document exists.
	filter := bson.M{"ticker": transaction.Ticker}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}

	if count == 0 {
		newStockTransaction := StockTransactions{
			Ticker:       transaction.Ticker,
			Transactions: []Transaction{transaction},
		}

		// Insert the new document.
		return c.upsertStockTransaction(ctx, filter, bson.M{"$set": newStockTransaction})
	}

	// If the document exists, push the new transaction into the array.
	// If the transaction already exists, it will not be added again.
	update := bson.M{
		"$addToSet": bson.M{"transactions": transaction},
	}

	return c.upsertStockTransaction(ctx, filter, update)
}
