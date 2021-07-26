package Required

import (
	"fmt"
	"testing"
)

func TestAppendNotNil(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		output *[]string
	}{
		{
			name:   "CaseNotNil",
			input:  "value",
			output: &[]string{"value"},
		},
		{
			name:   "CaseNil",
			input:  "value",
			output: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var missingFields *[]string
			if tt.name == "CaseNil" {
				missingFields = nil
			}
			AppendNotNil(missingFields, tt.input)
			if tt.name == "CaseNil" {
				if fmt.Sprint(nil) != fmt.Sprint(tt.output) {
					t.Errorf(
						"expected value on AppendNotNil(%v,%v) = %v; got %v",
						nil,
						tt.input,
						tt.output,
						nil,
					)
				}
			} else {
				if fmt.Sprint(&missingFields) != fmt.Sprint(tt.output) {
					t.Errorf(
						"expected value on AppendNotNil(%v,%v) = %v; got %v",
						&[]string{},
						tt.input,
						tt.output,
						missingFields,
					)
				}
			}

		})
	}
}
