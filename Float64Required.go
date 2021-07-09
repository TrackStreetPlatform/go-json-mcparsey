package go_json_mcparsey

import "strconv"

func Float64Required(origin map[string]interface{}, key string, requiredFields *[]string, defaultValue float64) (value float64, isValid bool) {
	if maybePrice, ok := origin[key]; ok {
		switch tempPrice := maybePrice.(type) {
		case string:
			Price, err := strconv.ParseFloat(tempPrice, 64)
			if err != nil {
				AppendWhenNotNil(requiredFields, key)
				return defaultValue, false
			}
			return Price, true
		case int:
			return float64(tempPrice), true
		case float64:
			return tempPrice, true
		default:
			AppendWhenNotNil(requiredFields, key)
		}
	} else {
		AppendWhenNotNil(requiredFields, key)
	}
	return defaultValue, false
}
