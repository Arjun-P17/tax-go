package testutils

import (
	"encoding/json"
	"os"

	"github.com/Arjun-P17/tax-go/internal/repository"
)

func SaveToFile(transactions []*repository.Transaction, outputfile string) error {
	data, err := json.Marshal(transactions)
	if err != nil {
		return err
	}
	err = os.WriteFile(outputfile, data, 0o644)
	if err != nil {
		return err
	}
	return nil
}

func LoadTransactionsFromFile(filename string) ([]*repository.Transaction, error) {
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
