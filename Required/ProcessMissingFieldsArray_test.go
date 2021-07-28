package Required

import (
	"errors"
	"fmt"
	"testing"
)

func TestProcessMissingFieldsArray(t *testing.T) {
	type inputStruct struct {
		MissingFields []string
		StrType       string
	}
	tests := []struct {
		name   string
		input  inputStruct
		output error
	}{
		{
			name: "EmptyMissingFields",
			input: inputStruct{
				MissingFields: []string{},
				StrType:       "test",
			},
			output: nil,
		},
		{
			name: "NonEmptyMissingFields",
			input: inputStruct{
				MissingFields: []string{"value", "attribute"},
				StrType:       "test",
			},
			output: errors.New("required fields in json not available in the correct type for  test :  value, attribute"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			errGot := ProcessMissingFieldsArray(tt.input.MissingFields, tt.input.StrType)
			if fmt.Sprint(errGot) != fmt.Sprint(tt.output) {
				t.Errorf(
					"expected ProcessMissingFieldsArray(%v,%v) = %v; got %v",
					tt.input.MissingFields,
					tt.input.StrType,
					tt.output,
					errGot,
				)
			}
		})
	}
}
