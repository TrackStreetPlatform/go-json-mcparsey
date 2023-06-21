package Optional

import (
	"github.com/TrackStreetPlatform/go-json-mcparsey/Path"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UUID(origin map[string]interface{}, path string, defaultValue uuid.UUID) uuid.UUID {
	maybeValueInField, err := Path.Traverse(origin, path)
	if err == nil {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			parse, err := uuid.Parse(tempValueInField)
			if err != nil {
				return defaultValue
			}
			return parse
		case primitive.Binary:
			parse, err := uuid.FromBytes(tempValueInField.Data)
			if err != nil {
				return defaultValue
			}
			return parse
		case []byte:
			parse, err := uuid.FromBytes(tempValueInField)
			if err != nil {
				return defaultValue
			}
			return parse
		case uuid.UUID:
			return tempValueInField
		default:
			break
		}
	}
	return defaultValue
}
