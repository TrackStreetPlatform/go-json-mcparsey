package JSONHelper

import (
	"errors"
	"fmt"
	"strings"
)

func ProcessRequiredFieldsArray(requiredFields []string, strType string) error {
	if len(requiredFields) > 0 {
		fields := strings.Join(requiredFields, ", ")
		errorMsg := fmt.Sprintf("%s %s ", "required fields in json not avaiable in the correct type for ", strType, ": ", fields)
		return errors.New(errorMsg)
	}
	return nil
}
