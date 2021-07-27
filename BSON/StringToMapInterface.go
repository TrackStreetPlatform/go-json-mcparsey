package JSON

import (
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func StringToMapInterface(messageBody []byte, typeStr string) (map[string]interface{}, error) {
	var temp interface{}
	err := bson.Unmarshal(messageBody, &temp)
	if err != nil {
		return nil, err
	}
	switch tempResponse := temp.(type) {
	case map[string]interface{}:
		return tempResponse, nil
	default:
		return nil, errors.New(fmt.Sprint("Bson not in format expected(", typeStr, "), got format ", tempResponse, ""))
	}
}
