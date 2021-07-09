package JSONHelper

func StringOptional(origin map[string]interface{}, key string, defaultValue string) string {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			return tempValueInField
		default:
			break
		}
	}
	return defaultValue
}
