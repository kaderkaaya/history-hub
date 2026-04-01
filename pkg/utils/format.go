package utils

import "fmt"

func NormalizeMonthDay(value string) string {
	var num int
	fmt.Sscanf(value, "%d", &num)
	return fmt.Sprintf("%02d", num)
}
