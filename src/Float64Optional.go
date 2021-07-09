package JSONHelper

import "strconv"

func Float64Optional(origin map[string]interface{}, key string, defaultValue float64) float64 {
	if maybePrice, ok := origin[key]; ok {
		switch tempPrice := maybePrice.(type) {
		case string:
			Price, err := strconv.ParseFloat(tempPrice, 64)
			if err != nil {
				return defaultValue
			}
			return Price
		case int:
			return float64(tempPrice)
		case float64:
			return tempPrice
		default:
			break
		}
	}
	return defaultValue
}
