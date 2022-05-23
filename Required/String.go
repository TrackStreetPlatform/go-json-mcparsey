package Required

import "github.com/TrackStreetPlatform/go-json-mcparsey/Path"

func String(origin map[string]interface{}, path string, missingFields *[]string, defaultValue string) (value string, isValid bool) {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return tempValueInField, true
		default:
			AppendNotNil(missingFields, path)
		}
	} else {
		AppendNotNil(missingFields, path)
	}
	return defaultValue, false
}
