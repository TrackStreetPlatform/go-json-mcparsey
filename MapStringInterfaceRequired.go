package go_json_mcparsey

func MapStringInterfaceRequired(origin map[string]interface{}, key string, requiredFields *[]string, defaultValue map[string]interface{}) (value map[string]interface{}, isValid bool) {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case map[string]interface{}:
			return tempValueInField, true
		default:
			AppendWhenNotNil(requiredFields, key)
		}
	} else {
		AppendWhenNotNil(requiredFields, key)
	}
	return defaultValue, false
}
