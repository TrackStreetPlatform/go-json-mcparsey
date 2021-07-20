package Optional

import "strings"

func ArrayString(origin map[string]interface{}, key string, defaultValue []string) []string {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return strings.Split(tempValueInField, ",")
		case []interface{}:
			items := []string{}
			for _, val := range tempValueInField {
				switch strItem := val.(type) {
				case string:
					items = append(items, strItem)
				default:
					return defaultValue
				}
			}
			return items
		case []string:
			return tempValueInField
		default:
			break
		}
	}
	return defaultValue
}
