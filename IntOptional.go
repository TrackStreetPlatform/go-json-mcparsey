package go_json_mcparsey

import "strconv"

func IntOptional(origin map[string]interface{}, key string, defaultValue int) int {
	if maybeProductId, ok := origin[key]; ok {
		switch tempProductId := maybeProductId.(type) {
		case string:
			ProductId, err := strconv.Atoi(tempProductId)
			if err != nil {
				return defaultValue
			}
			return ProductId
		case int:
			return tempProductId
		case int32:
			return int(tempProductId)
		case int64:
			return int(tempProductId)
		case float64:
			return int(tempProductId)
		default:
			break
		}
	}

	return defaultValue
}
