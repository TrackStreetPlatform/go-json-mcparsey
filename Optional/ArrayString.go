package Optional

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"strings"
)

func ArrayString(origin map[string]interface{}, path string, defaultValue []string) []string {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return strings.Split(tempValueInField, ",")
		case []interface{}:
			items := []string{}
			for _, val := range tempValueInField {
				switch strItem := val.(type) {
				case string:
					items = append(items, strItem)
				default:
					return defaultValue
				}
			}
			return items
		case []string:
			return tempValueInField
		default:
			break
		}
	}
	return defaultValue
}
