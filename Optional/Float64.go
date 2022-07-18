package Optional

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"strconv"
)

func Float64(origin map[string]interface{}, path string, defaultValue float64) float64 {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempVal := maybeValueInField.(type) {
		case string:
			Price, err := strconv.ParseFloat(tempVal, 64)
			if err != nil {
				return defaultValue
			}
			return Price
		case int:
			return float64(tempVal)
		case float64:
			return tempVal
		default:
			break
		}
	}
	return defaultValue
}
