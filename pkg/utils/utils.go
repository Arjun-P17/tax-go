package utils

import "time"

func StringToTime(dateString string) (time.Time, error) {
	dateFormat := "2006-01-02, 15:04:05"
	date, err := time.Parse(dateFormat, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

var tickerMap = map[string]string{"FB": "META"}

func GetMappedTicker(ticker string) string {
	mappedTicker := tickerMap[ticker]
	if mappedTicker == "" {
		mappedTicker = ticker
	}

	return mappedTicker
}
