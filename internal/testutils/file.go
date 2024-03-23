package testutils

import (
	"encoding/json"
	"os"

	"github.com/Arjun-P17/tax-go/internal/app/parseTrades/service"
	"github.com/Arjun-P17/tax-go/internal/repository"
)

func SaveToFile(inputcsv string, outputfile string) error {
	file, err := os.Open(inputcsv)
	if err != nil {
		return err
	}
	defer file.Close()

	transactions, err := service.ParseTransactions(nil)
	if err != nil {
		return err
	}

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

func LoadFromFile(filename string) ([]*repository.Transaction, error) {
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
