package Required

func MapStringInterface(origin map[string]interface{}, key string, missingFields *[]string, defaultValue map[string]interface{}) (value map[string]interface{}, isValid bool) {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case map[string]interface{}:
			return tempValueInField, true
		default:
			AppendNotNil(missingFields, key)
		}
	} else {
		AppendNotNil(missingFields, key)
	}
	return defaultValue, false
}
