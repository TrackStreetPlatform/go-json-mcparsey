package Required

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"strconv"
)

func Float64(origin map[string]interface{}, path string, missingFields *[]string, defaultValue float64) (value float64, isValid bool) {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValue := maybeValueInField.(type) {
		case string:
			value, err := strconv.ParseFloat(tempValue, 64)
			if err != nil {
				AppendNotNil(missingFields, path)
				return defaultValue, false
			}
			return value, true
		case int:
			return float64(tempValue), true
		case float64:
			return tempValue, true
		default:
			AppendNotNil(missingFields, path)
		}
	} else {
		AppendNotNil(missingFields, path)
	}
	return defaultValue, false
}
