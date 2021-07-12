package Required

import "strconv"

func Bool(origin map[string]interface{}, key string, requiredFields *[]string, defaultValue bool) (value bool, isValid bool) {
	if maybeForce, ok := origin[key]; ok {
		switch tempForce := maybeForce.(type) {
		case string:
			boolValue, err := strconv.ParseBool(tempForce)
			if err != nil {
				AppendWhenNotNil(requiredFields, key)
				return false, defaultValue
			}
			return boolValue, true
		case int:
			return tempForce != 0, true
		case float64:
			return tempForce != 0, true
		case bool:
			return tempForce, true
		default:
			AppendWhenNotNil(requiredFields, key)
			break
		}
	} else {
		AppendWhenNotNil(requiredFields, key)
	}
	return defaultValue, false
}
