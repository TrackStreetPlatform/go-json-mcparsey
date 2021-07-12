package Required

func String(origin map[string]interface{}, key string, requiredFields *[]string, defaultValue string) (value string, isValid bool) {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return tempValueInField, true
		default:
			AppendWhenNotNil(requiredFields, key)
		}
	} else {
		AppendWhenNotNil(requiredFields, key)
	}
	return defaultValue, false
}
