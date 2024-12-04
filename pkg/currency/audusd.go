package currency

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/Arjun-P17/tax-go/pkg/date"
)

const currencyDateFormat = "2/01/2006"

func GetAUDUSDMap() (map[date.Date]float64, error) {
	rows, err := readAUDUSDCSV()
	if err != nil {
		return make(map[date.Date]float64), err
	}

	currencyMap, err := createCurrencyMap(rows)
	if err != nil {
		return make(map[date.Date]float64), err
	}

	return currencyMap, nil

}

func readAUDUSDCSV() ([][]string, error) {
	// Open the CSV file
	file, err := os.Open("AUD_USD_Historical_Data.csv")
	if err != nil {
		fmt.Printf("Error opening file: %v\n", err)
		return nil, errors.New("error opening file")
	}
	defer file.Close()

	// Read the CSV data
	reader := csv.NewReader(file)
	rows, err := reader.ReadAll()
	if err != nil {
		fmt.Printf("Error reading CSV: %v\n", err)
		return nil, errors.New("error reading file")
	}

	return rows, nil
}

func createCurrencyMap(rows [][]string) (map[date.Date]float64, error) {
	currencyMap := make(map[date.Date]float64)

	// Parse the rows into a map
	for i := 0; i < len(rows); i++ {
		row := rows[i]
		if i == 0 {
			// Skip header row
			continue
		}
		parsedDate, err := time.Parse(currencyDateFormat, row[0])
		if err != nil {
			errorMsg := fmt.Sprintf("Error parsing date on row %d: %v", i+1, err)
			fmt.Println(errorMsg)
			return currencyMap, errors.New(errorMsg)
		}

		price, err := strconv.ParseFloat(row[1], 64)
		if err != nil {
			errorMsg := fmt.Sprintf("Error parsing price on row %d: %v", i+1, err)
			fmt.Println(errorMsg)
			return currencyMap, errors.New(errorMsg)
		}

		currencyMap[date.NewDate(parsedDate)] = price

		// If all rows processed break to prevent out of bounds access
		if i == len(rows)-1 {
			break
		}

		nextRow := rows[i+1]
		nextDate, err := time.Parse(currencyDateFormat, nextRow[0])
		if err != nil {
			errorMsg := fmt.Sprintf("Error parsing date on row %d: %v", i+1, err)
			fmt.Println(errorMsg)
			return currencyMap, errors.New(errorMsg)
		}

		currDate := parsedDate
		for nextDate.Sub(currDate) > 24*time.Hour {
			currDate = currDate.Add(24 * time.Hour)
			currencyMap[date.NewDate(currDate)] = price
		}

	}

	return currencyMap, nil
}
