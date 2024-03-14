package service

import (
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"testing"

	"github.com/Arjun-P17/tax-go/internal/repository"
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

	loadedTransactions, err := loadFromFile("./testdata/sample.txt")
	assert.Nil(t, err)

	if reflect.DeepEqual(parsedTransactions, loadedTransactions) {
		fmt.Println("Loaded slice matches original slice")
	} else {
		fmt.Println("Loaded slice does not match original slice")
	}

}

// func saveToFile(transactions []*repository.Transaction, filename string) error {
// 	data, err := json.Marshal(transactions)
// 	if err != nil {
// 		return err
// 	}
// 	err = os.WriteFile(filename, data, 0644)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func loadFromFile(filename string) ([]*repository.Transaction, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	var transactions []*repository.Transaction
	err = json.Unmarshal(data, &transactions)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}
