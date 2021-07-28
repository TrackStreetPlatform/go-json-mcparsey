package Required

import (
	"errors"
	"fmt"
	"strings"
)

func ProcessMissingFieldsArray(missingFields []string, strType string) error {
	if len(missingFields) > 0 {
		fields := strings.Join(missingFields, ", ")
		errorMsg := fmt.Sprintf("%s %s %s %s", "required fields in json not available in the correct type for ", strType, ": ", fields)
		return errors.New(errorMsg)
	}
	return nil
}
