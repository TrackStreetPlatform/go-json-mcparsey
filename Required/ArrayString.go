package Required

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"strings"
)

func ArrayString(origin map[string]interface{}, path string, missingFields *[]string, defaultValue []string) (value []string, isValid bool) {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return strings.Split(tempValueInField, ","), true
		case []interface{}:
			var items []string
			for _, val := range tempValueInField {
				switch strItem := val.(type) {
				case string:
					items = append(items, strItem)
				default:
					AppendNotNil(missingFields, path)
					return defaultValue, false
				}
			}
			return items, true
		case []string:
			return tempValueInField, true
		default:
			break
		}
	}
	AppendNotNil(missingFields, path)
	return defaultValue, false
}
