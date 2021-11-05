package Required

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func Date(origin map[string]interface{}, key string, missingFields *[]string, defaultValue time.Time) (value time.Time, isValid bool) {
	var err error
	if maybeForce, ok := origin[key]; ok {
		switch tempForce := maybeForce.(type) {
		case string:
			layouts := []string{
				"2006-01-02 15:04:05",
				"2006-01-02T15:04:05Z",
				"2006-01-02T15:04:05.999Z",
				"2006-01-02T15:04:05.999999999Z",
			}
			for _, layout := range layouts {
				value, err = time.Parse(layout, tempForce)
				if err == nil {
					return value, true
				}
			}
			if err != nil {
				AppendNotNil(missingFields, key)
				return defaultValue, false
			}
		case time.Time:
			return tempForce, true
		case primitive.DateTime:
			return tempForce.Time(), true
		default:
			AppendNotNil(missingFields, key)
			break
		}
	} else {
		AppendNotNil(missingFields, key)
	}
	return defaultValue, false
}
