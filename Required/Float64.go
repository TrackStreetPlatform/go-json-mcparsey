package Required

import "strconv"

func Float64(origin map[string]interface{}, key string, missingFields *[]string, defaultValue float64) (value float64, isValid bool) {
	if maybeValue, ok := origin[key]; ok {
		switch tempValue := maybeValue.(type) {
		case string:
			value, err := strconv.ParseFloat(tempValue, 64)
			if err != nil {
				AppendNotNil(missingFields, key)
				return defaultValue, false
			}
			return value, true
		case int:
			return float64(tempValue), true
		case float64:
			return tempValue, true
		default:
			AppendNotNil(missingFields, key)
		}
	} else {
		AppendNotNil(missingFields, key)
	}
	return defaultValue, false
}
