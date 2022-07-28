package Required

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UUID(origin map[string]interface{}, path string, missingFields *[]string, defaultValue uuid.UUID) (value uuid.UUID, isValid bool) {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			parse, err := uuid.Parse(tempValueInField)
			if err != nil {
				AppendNotNil(missingFields, path)
				return defaultValue, false
			}
			return parse, true
		case primitive.Binary:
			parse, err := uuid.FromBytes(tempValueInField.Data)
			if err != nil {
				AppendNotNil(missingFields, path)
				return defaultValue, false
			}
			return parse, true
		case []byte:
			parse, err := uuid.FromBytes(tempValueInField)
			if err != nil {
				AppendNotNil(missingFields, path)
				return defaultValue, false
			}
			return parse, true
		case uuid.UUID:
			return tempValueInField, true
		default:
			break
		}
	}
	AppendNotNil(missingFields, path)
	return defaultValue, false
}
