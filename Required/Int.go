package Required

import "strconv"

func Int(origin map[string]interface{}, key string, missingFields *[]string, defaultValue int) (value int, isValid bool) {
	if maybeProductId, ok := origin[key]; ok {
		switch tempProductId := maybeProductId.(type) {
		case string:
			ProductId, err := strconv.Atoi(tempProductId)
			if err != nil {
				AppendNotNil(missingFields, key)
				return defaultValue, false
			}
			return ProductId, true
		case int:
			return tempProductId, true
		case int32:
			return int(tempProductId), true
		case int64:
			return int(tempProductId), true
		case float64:
			return int(tempProductId), true
		default:
			AppendNotNil(missingFields, key)
		}
	} else {
		AppendNotNil(missingFields, key)
	}

	return defaultValue, false
}
