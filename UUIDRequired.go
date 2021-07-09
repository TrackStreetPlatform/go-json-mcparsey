package go_json_mcparsey

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	uuid2 "go.mongodb.org/mongo-driver/x/mongo/driver/uuid"
)

func UUIDRequired(origin map[string]interface{}, key string, requiredFields *[]string, defaultValue uuid.UUID) (value uuid.UUID, isValid bool) {
	if maybeValueInField, ok := origin[key]; ok {
		switch tempValueInField := maybeValueInField.(type) {
		case string:
			parse, err := uuid.Parse(tempValueInField)
			if err != nil {
				AppendWhenNotNil(requiredFields, key)
				return defaultValue, false
			}
			return parse, true
		case primitive.Binary:
			parse, err := uuid.FromBytes(tempValueInField.Data)
			if err != nil {
				AppendWhenNotNil(requiredFields, key)
				return defaultValue, false
			}
			return parse, true
		case []byte:
			parse, err := uuid.FromBytes(tempValueInField)
			if err != nil {
				AppendWhenNotNil(requiredFields, key)
				return defaultValue, false
			}
			return parse, true
		case uuid.UUID:
			return tempValueInField, true
		case uuid2.UUID:
			parse, err := uuid.FromBytes(tempValueInField[:])
			if err != nil {
				AppendWhenNotNil(requiredFields, key)
				return defaultValue, false
			}
			return parse, true
		default:
			break
		}
	}
	AppendWhenNotNil(requiredFields, key)
	return defaultValue, false
}
