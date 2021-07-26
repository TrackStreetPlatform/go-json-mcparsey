package Required

func String(origin map[string]interface{}, key string, missingFields *[]string, defaultValue string) (value string, isValid bool) {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return tempValueInField, true
		default:
			AppendNotNil(missingFields, key)
		}
	} else {
		AppendNotNil(missingFields, key)
	}
	return defaultValue, false
}
