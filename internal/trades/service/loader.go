package service

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/Arjun-P17/tax-go/models"
	"github.com/Arjun-P17/tax-go/pkg/utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// ParseTransactions reads a CSV file and creates a list of Transaction objects.
func ParseTransactions(csvFilePath string) ([]*models.Transaction, error) {
	transactions := make([]*models.Transaction, 0)

	file, err := os.Open(csvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		transaction, err := parseTransaction(row)
		if err != nil {
			return nil, err
		}
		if transaction != nil {
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}

// parseTransaction parses a CSV row and creates a Transaction object.
func parseTransaction(row []string) (*models.Transaction, error) {
	if len(row) < 4 {
		fmt.Println("transaction not in correct format continuing")
		return nil, nil
	}

	fmt.Println(row)

	key := row[0]
	order := row[2]
	asset := row[3]
	if key != "Trades" || asset != "Stocks" || order != "Order" {
		fmt.Println("transaction not in correct format continuing")
		return nil, nil
	}

	currency := row[4]
	ticker := utils.GetMappedTicker(row[5])
	date, err := utils.StringToTime(row[6])
	if err != nil {
		return nil, err
	}
	// positive if a buy
	quantity, err := strconv.ParseFloat(row[7], 64)
	if err != nil {
		return nil, err
	}
	tradePrice, err := strconv.ParseFloat(row[8], 64)
	if err != nil {
		return nil, err
	}
	// negative if a buy
	proceeds, err := strconv.ParseFloat(row[10], 64)
	if err != nil {
		return nil, err
	}
	// always negative
	brokerageFee, err := strconv.ParseFloat(row[11], 64)
	if err != nil {
		return nil, err
	}
	brokerageFee *= -1
	basis, err := strconv.ParseFloat(row[12], 64)
	if err != nil {
		return nil, err
	}
	brokerProfit, err := strconv.ParseFloat(row[13], 64)
	if err != nil {
		return nil, err
	}

	typ := models.Buytype
	if quantity < 0 {
		typ = models.Selltype
		quantity *= -1
		basis *= -1
		brokerProfit *= -1
	} else {
		proceeds *= -1
	}

	// Include broker fees in the proceeds of trade
	proceeds += brokerageFee
	realPrice := proceeds / quantity

	// Handle stock splits
	splitFactor, err := getSplitFactor(ticker, date)
	if err != nil {
		fmt.Println(ticker, err)
		return nil, err
	}
	quantity, tradePrice, realPrice = quantity*splitFactor, tradePrice/splitFactor, realPrice/tradePrice

	return &models.Transaction{
		ID:           primitive.NewObjectID(),
		Ticker:       ticker,
		Currency:     currency,
		Date:         date,
		Type:         typ,
		Quantity:     quantity,
		TradePrice:   tradePrice,
		RealPrice:    realPrice,
		Proceeds:     proceeds,
		BrokerageFee: brokerageFee,
		Basis:        basis,
		BrokerProfit: brokerProfit,
		USDAUD:       1.45,
		Splitfactor:  splitFactor,
	}, nil
}

var splitFactors = map[string]struct {
	Date   string
	Factor float64
}{
	"AMZN": {"2022-06-06, 00:00:00", 20.0},
	"SHOP": {"2022-06-29, 00:00:00", 10.0},
	"TSLA": {"2022-08-25, 00:00:00", 3.0},
	"PANW": {"2022-09-14, 00:00:00", 3.0},
}

// getSplitfactor returns the splitfactor for a stock split, if transaction date is before splitdate so we can normalise the transaction values relative to post split
func getSplitFactor(ticker string, date time.Time) (float64, error) {
	info, exists := splitFactors[ticker]
	if !exists {
		return 1.0, nil
	}

	splitDateParsed, err := utils.StringToTime(info.Date)
	if err != nil {
		return 1.0, err
	}

	if date.After(splitDateParsed) {
		return 1.0, nil
	}

	return info.Factor, nil
}
