package Optional

import "strings"

func ArrayString(origin map[string]interface{}, key string, defaultValue []string) []string {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return strings.Split(tempValueInField, ",")
		default:
			break
		}
	}
	return defaultValue
}
