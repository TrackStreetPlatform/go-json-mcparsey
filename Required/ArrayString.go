package Required

import "strings"

func ArrayString(origin map[string]interface{}, key string, requiredFields *[]string, defaultValue []string) (value []string, isValid bool) {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return strings.Split(tempValueInField, ","), false
		default:
			AppendWhenNotNil(requiredFields, key)
			break
		}
	}
	return defaultValue, false
}
