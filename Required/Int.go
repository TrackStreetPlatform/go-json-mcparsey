package Required

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"strconv"
)

func Int(origin map[string]interface{}, path string, missingFields *[]string, defaultValue int) (value int, isValid bool) {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tmpVal := maybeValueInField.(type) {
		case string:
			ProductId, err := strconv.Atoi(tmpVal)
			if err != nil {
				AppendNotNil(missingFields, path)
				return defaultValue, false
			}
			return ProductId, true
		case int:
			return tmpVal, true
		case int32:
			return int(tmpVal), true
		case int64:
			return int(tmpVal), true
		case float64:
			return int(tmpVal), true
		default:
			AppendNotNil(missingFields, path)
		}
	} else {
		AppendNotNil(missingFields, path)
	}

	return defaultValue, false
}
