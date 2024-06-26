package repository

import (
	"context"

	"github.com/Arjun-P17/tax-go/pkg/utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (c *Repository) upsertStockTax(ctx context.Context, filter bson.M, update bson.M) error {
	collection := c.GetCollection(c.config.TaxCollection)

	options := options.Update().SetUpsert(true)
	_, err := collection.UpdateOne(ctx, filter, update, options)
	return err
}

// TODO: should this be an upsert like UpsertStockPosition and let the service layer handle the logic to update values?
// TODO: should the buy amount be coverted to aud when calculating profit?
func (c *Repository) InsertTaxEvent(ctx context.Context, taxEvent TaxEvent, USDAUD float64) error {
	collection := c.GetCollection(c.config.TaxCollection)

	finYear := utils.GetFYYearString(taxEvent.Date)
	// Check if the document exists.
	filter := bson.M{"financialyear": finYear}
	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return err
	}

	var gainsCGT, gains, losses float64 = 0, 0, 0
	if taxEvent.Profit > 0 {
		gainsCGT = taxEvent.CGTProfit
		gains = taxEvent.Profit
	} else {
		losses = taxEvent.Profit
	}

	if count == 0 {
		// If the document does not exist, create and insert new document.
		newStockTax := StockTax{
			FinancialYear:   finYear,
			NetProfitCGT:    taxEvent.CGTProfit,
			NetProfitCGTAUD: taxEvent.CGTProfit * USDAUD,
			NetProfit:       taxEvent.Profit,
			NetProfitAUD:    taxEvent.Profit * USDAUD,
			GainsCGT:        gainsCGT,
			GainsCGTAUD:     gainsCGT * USDAUD,
			Gains:           gains,
			Losses:          losses,
			Events:          []TaxEvent{taxEvent},
		}
		return c.upsertStockTax(ctx, filter, bson.M{"$set": newStockTax})
	}

	// If the document exists, push the new transaction into the array and update other values.
	update := bson.M{
		"$push": bson.M{"events": taxEvent},
		"$inc": bson.M{
			"netprofitcgt":    taxEvent.CGTProfit,
			"netprofitcgtaud": taxEvent.CGTProfit * USDAUD,
			"netprofit":       taxEvent.Profit,
			"netprofitaud":    taxEvent.Profit * USDAUD,
			"gainscgt":        gainsCGT,
			"gainscgtaud":     gainsCGT * USDAUD,
			"gains":           gains,
			"losses":          losses,
		},
	}
	return c.upsertStockTax(ctx, filter, update)
}
