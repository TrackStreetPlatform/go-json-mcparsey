package go_json_mcparsey

import (
	"encoding/json"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
)

func JSONStringToMapInterface(messageBody string, typeStr string) (map[string]interface{}, error) {
	var temp interface{}
	err := json.Unmarshal([]byte(messageBody), &temp)
	if err != nil {
		return nil, err
	}
	var maybeTask map[string]interface{}
	switch tempResponse := temp.(type) {
	case map[string]interface{}:
		maybeTask = tempResponse
		return maybeTask, nil
	default:
		return nil, errors.New(fmt.Sprint("Json not in format expected(", typeStr, ")"))
	}
}

func BSONStringToMapInterface(messageBody []byte, typeStr string) (map[string]interface{}, error) {
	var temp interface{}
	err := bson.Unmarshal(messageBody, &temp)
	if err != nil {
		return nil, err
	}
	var maybeTask map[string]interface{}
	switch tempResponse := temp.(type) {
	case map[string]interface{}:
		maybeTask = tempResponse
		return maybeTask, nil
	default:
		return nil, errors.New(fmt.Sprint("Bson not in format expected(", typeStr, ")"))
	}
}
