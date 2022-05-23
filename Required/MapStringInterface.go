package Required

import "github.com/TrackStreetPlatform/go-json-mcparsey/Path"

func MapStringInterface(origin map[string]interface{}, path string, missingFields *[]string, defaultValue map[string]interface{}) (value map[string]interface{}, isValid bool) {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case map[string]interface{}:
			return tempValueInField, true
		default:
			AppendNotNil(missingFields, path)
		}
	} else {
		AppendNotNil(missingFields, path)
	}
	return defaultValue, false
}
