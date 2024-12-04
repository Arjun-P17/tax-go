package service

import (
	"encoding/csv"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/Arjun-P17/tax-go/pkg/utils"
	"github.com/Arjun-P17/tax-go/pkg/utils/slices"
)

// ParseTransactions reads a CSV file and creates a list of Transaction objects.
func ParseTransactions(file *os.File) ([]*repository.Transaction, error) {
	if file == nil {
		return nil, errors.New("nil file")
	}

	transactions := make([]*repository.Transaction, 0)

	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		transaction, err := parseIBKRTransaction(row)
		if err != nil {
			return nil, err
		}
		if transaction != nil {
			transactions = append(transactions, transaction)
		}
	}

	return transactions, nil
}

// parseIBKRTransaction parses a CSV row and creates a Transaction object. Prices are adjusted to reflect current stock splits
func parseIBKRTransaction(row []string) (*repository.Transaction, error) {
	if len(row) < 4 {
		fmt.Println("transaction not in correct format continuing")
		return nil, nil
	}

	fmt.Println(row)

	key := row[0]
	order := row[2]
	asset := row[3]

	allowableAssets := []string{"Stocks", "Equity and Index Options"}
	if key != "Trades" || order != "Order" || !slices.Contains(allowableAssets, asset) {
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

	// Include broker fees in the proceeds of trade
	typ := repository.Buytype
	if quantity < 0 {
		typ = repository.Selltype
		quantity *= -1
		basis *= -1
		proceeds -= brokerageFee
	} else {
		proceeds *= -1
		proceeds += brokerageFee
	}
	// real price is the price per unit after brokerage fees
	realPrice := proceeds / quantity

	fmt.Println(proceeds, quantity, realPrice)

	// Handle stock splits
	splitFactor, err := utils.GetSplitFactor(ticker, date)
	if err != nil {
		fmt.Println(ticker, err)
		return nil, err
	}
	// adjust quantity and real price by the split factor
	quantity, realPrice = quantity*splitFactor, realPrice/splitFactor

	id := fmt.Sprintf("%s_%s_%v_%v", utils.TimeToString(date), ticker, quantity, tradePrice)
	return &repository.Transaction{
		ID:           id,
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
