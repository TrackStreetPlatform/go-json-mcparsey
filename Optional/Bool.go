package Optional

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"strconv"
)

func Bool(origin map[string]interface{}, path string, defaultValue bool) bool {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tmpVal := maybeValueInField.(type) {
		case string:
			Valid, err := strconv.ParseBool(tmpVal)
			if err != nil {
				return defaultValue
			}
			return Valid
		case int:
			return tmpVal != 0
		case float64:
			return tmpVal != 0
		case bool:
			return tmpVal
		default:
			break
		}
	}
	return defaultValue
}
