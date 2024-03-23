package service

import (
	"testing"

	"github.com/Arjun-P17/tax-go/internal/repository"
	"github.com/stretchr/testify/assert"
)

func Test_processTransaction(t *testing.T) {
	allTransactions, _, _, err := LoadTransactionsFromCSV("./testdata/fifo1.csv")
	assert.Nil(t, err)

	stockPosition := repository.StockPosition{}

	for _, transaction := range allTransactions {
		unique := processTransaction(&stockPosition, transaction)

		assert.True(t, unique)
	}
}
