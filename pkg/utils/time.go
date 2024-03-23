package utils

import (
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
