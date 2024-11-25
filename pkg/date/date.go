package date

import "time"

type Date struct {
	date time.Time
}

func NewDate(date time.Time) Date {
	zeroTimeDate := time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())
	return Date{
		date: zeroTimeDate,
	}
}

func (d Date) DateToString(t time.Time) string {
	return t.Format("2006-01-02")
}
