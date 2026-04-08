package cache

import "fmt"

func BuildOnThisDayKey(lang, typ, month, day string) string {
	return fmt.Sprintf("historyhub:onthisday:%s:%s:%s:%s", lang, typ, month, day)
}
