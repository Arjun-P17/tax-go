package service

import (
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/Arjun-P17/tax-go/internal/testutils"

	"github.com/stretchr/testify/assert"
)

func TestParseTransactions(t *testing.T) {
	file, err := os.Open("./testdata/sample.csv")
	if err != nil {
		assert.Nil(t, err)
	}
	defer file.Close()

	parsedTransactions, err := ParseTransactions(file)
	assert.Nil(t, err)

	loadedTransactions, err := testutils.LoadFromFile("./testdata/sample.txt")
	assert.Nil(t, err)

	if reflect.DeepEqual(parsedTransactions, loadedTransactions) {
		fmt.Println("Loaded slice matches original slice")
	} else {
		fmt.Println("Loaded slice does not match original slice")
	}
}
