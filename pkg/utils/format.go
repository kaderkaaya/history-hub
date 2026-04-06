package utils

import "fmt"

func NormalizeMonthDay(value string) string {
	var num int
	fmt.Sscanf(value, "%d", &num)
	return fmt.Sprintf("%02d", num)
}

func NormalizeMonthDayInt(value int) string {
	return fmt.Sprintf("%02d", value)
}
