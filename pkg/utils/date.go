package utils

import (
	"fmt"
	"time"
)

func ValidateMonthDay(month, day string) error {
	dateStr := fmt.Sprintf("2000-%s-%s", month, day)
	_, err := time.Parse("2006-01-02", dateStr)
	return err
}
