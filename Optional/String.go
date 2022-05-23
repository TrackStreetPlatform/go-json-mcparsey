package Optional

import "github.com/TrackStreetPlatform/go-json-mcparsey/Path"

func String(origin map[string]interface{}, path string, defaultValue string) string {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return tempValueInField
		default:
			break
		}
	}
	return defaultValue
}
