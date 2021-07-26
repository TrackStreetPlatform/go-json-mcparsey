package Required

func InterfaceArray(origin map[string]interface{}, key string, missingFields *[]string) (value []map[string]interface{}, isValid bool) {
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
			AppendNotNil(missingFields, key)
		}
	} else {
		AppendNotNil(missingFields, key)
	}
	return items, isValid
}
