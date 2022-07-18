package Optional

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func Date(origin map[string]interface{}, path string, defaultValue time.Time) time.Time {
	var err error
	var value time.Time
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tmpVal := maybeValueInField.(type) {
		case string:
			layouts := []string{
				"2006-01-02 15:04:05",
				"2006-01-02T15:04:05Z",
				"2006-01-02T15:04:05.999Z",
				"2006-01-02T15:04:05.999999999Z",
			}
			for _, layout := range layouts {
				value, err = time.Parse(layout, tmpVal)
				if err == nil {
					return value
				}
			}
			if err != nil {
				return defaultValue
			}
			return value
		case time.Time:
			return tmpVal
		case primitive.DateTime:
			return tmpVal.Time()
		default:
			break
		}
	}
	return defaultValue
}
