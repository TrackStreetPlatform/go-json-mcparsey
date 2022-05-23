package Required

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

func Date(origin map[string]interface{}, path string, missingFields *[]string, defaultValue time.Time) (value time.Time, isValid bool) {
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
					return value, true
				}
			}
			if err != nil {
				AppendNotNil(missingFields, path)
				return defaultValue, false
			}
		case time.Time:
			return tmpVal, true
		case primitive.DateTime:
			return tmpVal.Time(), true
		default:
			AppendNotNil(missingFields, path)
			break
		}
	} else {
		AppendNotNil(missingFields, path)
	}
	return defaultValue, false
}
