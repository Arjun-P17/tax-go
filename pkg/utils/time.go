package utils

import (
	"fmt"
	"time"
)

func StringToTime(dateString string) (time.Time, error) {
	dateFormat := "2006-01-02, 15:04:05"
	date, err := time.Parse(dateFormat, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}

func TimeToString(t time.Time) string {
	return t.Format("2006-01-02:15:04:05")
}

func IsOneYearGreaterThan(date1, date2 time.Time) bool {
	return date1.AddDate(1, 0, 0).Before(date2)
}

// GetFYYearString returns australian tax year in format YYYY-YYYY
func GetFYYearString(date time.Time) string {
	year, month, _ := date.Date()
	startYear := year
	if month < 7 {
		startYear = year - 1
	}
	endYear := startYear + 1
	taxYear := fmt.Sprintf("%d-%d", startYear, endYear)
	return taxYear
}
