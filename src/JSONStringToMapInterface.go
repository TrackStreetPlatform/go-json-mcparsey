package JSONHelper

import (
	"dullahan/extensions/CloudLogger"
	"encoding/json"
	"errors"
	"fmt"
)

func JSONStringToMapInterface(messageBody string, typeStr string) (map[string]interface{}, error) {
	var temp interface{}
	err := json.Unmarshal([]byte(messageBody), &temp)
	if err != nil {
		CloudLogger.Log(CloudLogger.ERROR, "json task Unmarshal not able to parse ", typeStr, " data", err, messageBody)
		return nil, err
	}
	var maybeTask map[string]interface{}
	switch tempResponse := temp.(type) {
	case map[string]interface{}:
		maybeTask = tempResponse
		return maybeTask, nil
	default:
		CloudLogger.Log(CloudLogger.ERROR, "json ", typeStr, " not in format expected", err, messageBody)
		return nil, errors.New(fmt.Sprint("Json not in format expected(", typeStr, ")"))
	}
}
