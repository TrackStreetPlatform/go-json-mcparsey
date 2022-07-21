package Optional

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"strconv"
	"strings"
)

func ArrayInt(origin map[string]interface{}, path string, defaultValue []int) []int {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			values := strings.Split(tempValueInField, ",")
			return mapStringToInt(values)
		case []interface{}:
			var items []int
			for _, val := range tempValueInField {
				switch arrItem := val.(type) {
				case string:
					intItem, err := strconv.Atoi(arrItem)
					if err != nil {
						items = append(items, intItem)
					}
				case int:
					items = append(items, arrItem)
				case int32:
					items = append(items, int(arrItem))
				case int64:
					items = append(items, int(arrItem))
				case float64:
					items = append(items, int(arrItem))
				default:
					return defaultValue
				}
			}
			return items
		case []string:
			return mapStringToInt(tempValueInField)
		case []int:
			return tempValueInField
		default:
			break
		}
	}
	return defaultValue
}

func mapStringToInt(values []string) []int {
	var items []int
	for _, arrItem := range values {
		intItem, err := strconv.Atoi(arrItem)
		if err != nil {
			items = append(items, intItem)
		}
	}
	return items
}
