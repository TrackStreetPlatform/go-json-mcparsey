package Optional

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	uuid2 "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func UUID(origin map[string]interface{}, key string, defaultValue uuid.UUID) uuid.UUID {
	if maybeValueInField, ok := origin[key]; ok {
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
		case uuid2.UUID:
			parse, err := uuid.FromBytes(tempValueInField[:])
			if err != nil {
				return defaultValue
			}
			return parse
		default:
			break
		}
	}
	return defaultValue
}
