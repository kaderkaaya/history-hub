package utils

var allowedTypes = map[string]struct{}{
	"events": {}, "births": {}, "deaths": {},
	"holidays": {}, "selected": {}, "all": {},
}

func IsValidType(typ string) bool {
	_, ok := allowedTypes[typ]
	return ok
}
