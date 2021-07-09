package JSONHelper

import "strconv"

func BoolOptional(origin map[string]interface{}, key string, defaultValue bool) bool {
	if maybeForce, ok := origin[key]; ok {
		switch tempForce := maybeForce.(type) {
		case string:
			Valid, err := strconv.ParseBool(tempForce)
			if err != nil {
				return defaultValue
			}
			return Valid
		case int:
			return tempForce != 0
		case float64:
			return tempForce != 0
		case bool:
			return tempForce
		default:
			break
		}
	}
	return defaultValue
}
