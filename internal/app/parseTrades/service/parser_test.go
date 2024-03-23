package service

import (
	"os"
	"testing"

	"github.com/Arjun-P17/tax-go/internal/testutils"

	"github.com/stretchr/testify/assert"
)

func TestParseTransactions(t *testing.T) {
	file, err := os.Open("./testdata/sample.csv")
	assert.Nil(t, err)
	defer file.Close()

	parsedTransactions, err := ParseTransactions(file)
	assert.Nil(t, err)

	loadedTransactions, err := testutils.LoadTransactionsFromFile("./testdata/sample.txt")
	assert.Nil(t, err)

	assert.Equal(t, parsedTransactions, loadedTransactions)
}
