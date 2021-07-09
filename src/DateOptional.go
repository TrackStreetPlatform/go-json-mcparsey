package JSONHelper

import (
	"time"
)

func DateOptional(origin map[string]interface{}, key string, defaultValue time.Time) time.Time {
	var err error
	var value time.Time
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
					return value
				}
			}
			if err != nil {
				return defaultValue
			}
			return value
		case time.Time:
			return tempForce
		default:
			break
		}
	}
	return defaultValue
}
