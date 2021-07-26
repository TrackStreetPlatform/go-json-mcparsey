package Required

import "strings"

func ArrayString(origin map[string]interface{}, key string, missingFields *[]string, defaultValue []string) (value []string, isValid bool) {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return strings.Split(tempValueInField, ","), true
		case []interface{}:
			var items []string
			for _, val := range tempValueInField {
				switch strItem := val.(type) {
				case string:
					items = append(items, strItem)
				default:
					AppendNotNil(missingFields, key)
					return defaultValue, false
				}
			}
			return items, true
		case []string:
			return tempValueInField, true
		default:
			break
		}
	}
	AppendNotNil(missingFields, key)
	return defaultValue, false
}
