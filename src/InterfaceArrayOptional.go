package JSONHelper

func InterfaceArrayOptional(origin map[string]interface{}, key string) []map[string]interface{} {
	items := make([]map[string]interface{}, 0)
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case []interface{}:
			for _, maybeItem := range tempValueInField {
				switch tempValueInField := maybeItem.(type) {
				case map[string]interface{}:
					items = append(items, tempValueInField)
				default:
					break
				}

			}
		default:
			break
		}
	}
	return items
}
