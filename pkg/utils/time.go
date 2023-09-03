package utils

import "time"

func BeginningOfMonth(date string) (string, error) {
	begDate, err := time.Parse(time.DateOnly, date+"-01")
	if err != nil {
		return "", err
	}
	return begDate.AddDate(0, 0, -begDate.Day()+1).Format(time.DateOnly), nil
}

func EndOfMonth(date string) (string, error) {
	endDate, err := time.Parse(time.DateOnly, date+"-01")
	if err != nil {
		return "", err
	}
	return endDate.AddDate(0, 1, -endDate.Day()).Format(time.DateOnly), nil
}

func GetMoscowTime() string {
	loc, _ := time.LoadLocation("Europe/Moscow")
	return time.Now().In(loc).Format(time.DateTime)
}
