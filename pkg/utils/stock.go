package utils

import "time"

/*
	Stock Split
*/

var splitFactors = map[string]struct {
	Date   string
	Factor float64
}{
	"AMZN": {"2022-06-03, 00:00:00", 20.0},
	"SHOP": {"2022-06-28, 00:00:00", 10.0},
	"TSLA": {"2022-08-24, 00:00:00", 3.0},
	"PANW": {"2022-09-14, 00:00:00", 3.0},
	"NVDA": {"2024-06-07, 00:00:00", 10.0},
	"MSTR": {"2024-08-07, 00:00:00", 10.0},
}

// getSplitfactor returns the splitfactor for a stock split, if transaction date is before splitdate so we can normalise the transaction values relative to post split
func GetSplitFactor(ticker string, date time.Time) (float64, error) {
	info, exists := splitFactors[ticker]
	if !exists {
		return 1.0, nil
	}

	splitDateParsed, err := StringToTime(info.Date)
	if err != nil {
		return 1.0, err
	}

	if date.After(splitDateParsed) {
		return 1.0, nil
	}

	return info.Factor, nil
}

/*
	Stock name changes
*/

var tickerMap = map[string]string{"FB": "META"}

func GetMappedTicker(ticker string) string {
	mappedTicker := tickerMap[ticker]
	if mappedTicker == "" {
		mappedTicker = ticker
	}

	return mappedTicker
}
