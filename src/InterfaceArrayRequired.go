package JSONHelper

func InterfaceArrayRequired(origin map[string]interface{}, key string, requiredFields *[]string) (value []map[string]interface{}, isValid bool) {
	items := make([]map[string]interface{}, 0)
	isValid = false
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case []interface{}:
			for _, maybeItem := range tempValueInField {
				switch tempValueInField := maybeItem.(type) {
				case map[string]interface{}:
					items = append(items, tempValueInField)
					isValid = true
				default:
					break
				}
			}
		default:
			AppendWhenNotNil(requiredFields, key)
		}
	} else {
		AppendWhenNotNil(requiredFields, key)
	}
	return items, isValid
}
