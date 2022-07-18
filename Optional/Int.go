package Optional

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"strconv"
)

func Int(origin map[string]interface{}, path string, defaultValue int) int {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tmpVal := maybeValueInField.(type) {
		case string:
			ProductId, err := strconv.Atoi(tmpVal)
			if err != nil {
				return defaultValue
			}
			return ProductId
		case int:
			return tmpVal
		case int32:
			return int(tmpVal)
		case int64:
			return int(tmpVal)
		case float64:
			return int(tmpVal)
		default:
			break
		}
	}

	return defaultValue
}
