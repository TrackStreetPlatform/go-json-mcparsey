package Optional

import "github.com/TrackStreetPlatform/go-json-mcparsey/Path"

func MapStringInterface(origin map[string]interface{}, path string, defaultValue map[string]interface{}) map[string]interface{} {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case map[string]interface{}:
			return tempValueInField
		default:
			break
		}
	}
	return defaultValue
}
