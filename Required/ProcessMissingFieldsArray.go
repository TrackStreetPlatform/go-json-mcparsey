package Required

import (
	"errors"
	"fmt"
	"strings"
)

func ProcessMissingFieldsArray(missingRequiredFields []string, strType string) error {
	if len(missingRequiredFields) > 0 {
		fields := strings.Join(missingRequiredFields, ", ")
		errorMsg := fmt.Sprintf("%s %s ", "required fields in json not avaiable in the correct type for ", strType, ": ", fields)
		return errors.New(errorMsg)
	}
	return nil
}
