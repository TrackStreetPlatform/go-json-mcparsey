package Required

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"strconv"
)

func Bool(origin map[string]interface{}, path string, missingFields *[]string, defaultValue bool) (value bool, isValid bool) {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tmpVal := maybeValueInField.(type) {
		case string:
			boolValue, err := strconv.ParseBool(tmpVal)
			if err != nil {
				AppendNotNil(missingFields, path)
				return defaultValue, false
			}
			return boolValue, true
		case int:
			return tmpVal != 0, true
		case float64:
			return tmpVal != 0, true
		case bool:
			return tmpVal, true
		default:
			AppendNotNil(missingFields, path)
			break
		}
	} else {
		AppendNotNil(missingFields, path)
	}
	return defaultValue, false
}
