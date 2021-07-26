package Required

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	uuid2 "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func UUID(origin map[string]interface{}, key string, missingFields *[]string, defaultValue uuid.UUID) (value uuid.UUID, isValid bool) {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			parse, err := uuid.Parse(tempValueInField)
			if err != nil {
				AppendNotNil(missingFields, key)
				return defaultValue, false
			}
			return parse, true
		case primitive.Binary:
			parse, err := uuid.FromBytes(tempValueInField.Data)
			if err != nil {
				AppendNotNil(missingFields, key)
				return defaultValue, false
			}
			return parse, true
		case []byte:
			parse, err := uuid.FromBytes(tempValueInField)
			if err != nil {
				AppendNotNil(missingFields, key)
				return defaultValue, false
			}
			return parse, true
		case uuid.UUID:
			return tempValueInField, true
		case uuid2.UUID:
			parse, err := uuid.FromBytes(tempValueInField[:])
			if err != nil {
				AppendNotNil(missingFields, key)
				return defaultValue, false
			}
			return parse, true
		default:
			break
		}
	}
	AppendNotNil(missingFields, key)
	return defaultValue, false
}
