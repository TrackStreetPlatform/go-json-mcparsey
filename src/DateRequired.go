package JSONHelper

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func DateRequired(origin map[string]interface{}, key string, requiredFields *[]string, defaultValue time.Time) (value time.Time, isValid bool) {
	var err error
	if maybeForce, ok := origin[key]; ok {
		switch tempForce := maybeForce.(type) {
		case string:
			layouts := []string{
				"2006-01-02 15:04:05",
				"2006-01-02T15:04:05Z",
			}
			for _, layout := range layouts {
				value, err = time.Parse(layout, tempForce)
				if err == nil {
					return value, true
				}
			}
			if err != nil {
				AppendWhenNotNil(requiredFields, key)
				return defaultValue, false
			}
		case primitive.DateTime:
			return tempForce.Time(), true
		default:
			AppendWhenNotNil(requiredFields, key)
			break
		}
	} else {
		AppendWhenNotNil(requiredFields, key)
	}
	return defaultValue, false
}
