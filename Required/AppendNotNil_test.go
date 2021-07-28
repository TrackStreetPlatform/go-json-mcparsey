package Required

import (
	"fmt"
	"testing"
)

func TestAppendNotNil(t *testing.T) {
	type inputStruct struct {
		MissingFields *[]string
		Key           string
	}
	tests := []struct {
		name   string
		input  inputStruct
		output *[]string
	}{
		{
			name: "CaseNotNil",
			input: inputStruct{
				MissingFields: &[]string{},
				Key:           "value",
			},
			output: &[]string{"value"},
		},
		{
			name: "CaseNil",
			input: inputStruct{
				MissingFields: nil,
				Key:           "value",
			},
			output: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AppendNotNil(tt.input.MissingFields, tt.input.Key)
			if tt.name == "CaseNil" {
				if fmt.Sprint(tt.input.MissingFields) != fmt.Sprint(tt.output) {
					t.Errorf(
						"expected value on AppendNotNil(%v,%v) = %v; got %v",
						nil,
						tt.input.Key,
						tt.output,
						tt.input.MissingFields,
					)
				}
			} else {
				if fmt.Sprint(tt.input.MissingFields) != fmt.Sprint(tt.output) {
					t.Errorf(
						"expected value on AppendNotNil(%v,%v) = %v; got %v",
						&[]string{},
						tt.input.Key,
						tt.output,
						tt.input.MissingFields,
					)
				}
			}
		})
	}
}
