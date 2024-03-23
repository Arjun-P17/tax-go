package service

import (
	"context"
	"os"
	"testing"

	"github.com/Arjun-P17/tax-go/internal/app/parseTrades/service"
	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/stretchr/testify/assert"
)

func Test_FIFO_1(t *testing.T) {
	buys, sells, err := LoadTransactionsFromCSV("./testdata/fifo1.csv")
	assert.Nil(t, err)

	output := fifo(context.Background(), sells[0], &buys)

	expectedOutput := ProcessSellOutput{
		Profit:    -6.882706435000025,
		CGTProfit: -6.882706435000025,
		BuysSold: []repository.BuySold{
			{
				Quantity: 1.0891,
				BuyID:    "2022-07-28:00:07:43_AAPL_1.0891_154.2483638",
			},
			{
				Quantity: 0.4023,
				BuyID:    "2022-09-13:00:13:13_AAPL_0.4023_161.5493",
			},
		},
	}

	assert.Equal(t, expectedOutput, output)
}

func Test_FIFO_2(t *testing.T) {
	buys, sells, err := LoadTransactionsFromCSV("./testdata/fifo2.csv")
	assert.Nil(t, err)

	tests := []ProcessSellOutput{
		{
			BuysSold: []repository.BuySold{{
				BuyID:    "2022-05-28:04:09:58_NVDA_1.3289_185.86505",
				Quantity: 0.7,
			}},
			Profit:    134.63731165326226,
			CGTProfit: 134.63731165326226,
		},
		{
			BuysSold: []repository.BuySold{{
				BuyID:    "2022-05-28:04:09:58_NVDA_1.3289_185.86505",
				Quantity: 0.4,
			}},
			Profit:    75.08950997700701,
			CGTProfit: 75.08950997700701,
		},
		{
			BuysSold: []repository.BuySold{
				{
					BuyID:    "2022-05-28:04:09:58_NVDA_1.3289_185.86505",
					Quantity: 0.2289,
				},
				{
					BuyID:    "2022-06-08:01:02:04_NVDA_1.5675_188.1927604",
					Quantity: 0.0711,
				},
			},
			Profit:    57.33589321691738,
			CGTProfit: 28.66794660845869,
		},
	}

	for index, test := range tests {
		output := fifo(context.Background(), sells[index], &buys)
		assert.Equal(t, test, output)
	}
}

func LoadTransactionsFromCSV(csvpath string) ([]repository.Buy, []repository.Transaction, error) {
	file, err := os.Open(csvpath)
	if err != nil {
		return nil, nil, err
	}
	defer file.Close()

	transactions, err := service.ParseTransactions(file)
	if err != nil {
		return nil, nil, err
	}

	buys := []repository.Buy{}
	sells := []repository.Transaction{}
	for _, transaction := range transactions {
		if transaction == nil {
			return nil, nil, err
		}

		if transaction.Type == repository.Buytype {
			buy := repository.Buy{
				Transaction:  *transaction,
				QuantityLeft: transaction.Quantity,
			}
			buys = append(buys, buy)
		} else {
			sells = append(sells, *transaction)
		}
	}

	return buys, sells, nil
}
