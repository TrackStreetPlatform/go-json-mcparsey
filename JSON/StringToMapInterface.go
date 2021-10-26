package JSON

import (
	"encoding/json"
	"errors"
	"fmt"
)

func StringToMapInterface(messageBody string, typeStr string) (map[string]interface{}, error) {
	var temp interface{}
	err := json.Unmarshal([]byte(messageBody), &temp)
	if err != nil {
		return nil, err
	}
	switch tempResponse := temp.(type) {
	case map[string]interface{}:
		return tempResponse, nil
	default:
		return nil, errors.New(fmt.Sprint("Json not in format expected(", typeStr, ")"))
	}
}

func StringToArrayInterface(messageBody string, typeStr string) ([]interface{}, error) {
	var temp interface{}
	err := json.Unmarshal([]byte(messageBody), &temp)
	if err != nil {
		return nil, err
	}
	switch tempResponse := temp.(type) {
	case []interface{}:
		return tempResponse, nil
	default:
		return nil, errors.New(fmt.Sprint("Json not in format expected(", typeStr, ")"))
	}
}
