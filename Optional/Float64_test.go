package Optional

import (
	"testing"
)

func TestFloat64(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue float64
	}
	tests := []struct {
		name   string
		input  inputStruct
		output float64
	}{
		{
			name: "NonExistingKey",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 42.},
				Key:          "NonExisting",
				DefaultValue: 0.,
			},
			output: 0.,
		},
		{
			name: "CaseInt",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 42},
				Key:          "value",
				DefaultValue: 0.,
			},
			output: 42.,
		},
		{
			name: "CaseFloat",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 42.},
				Key:          "value",
				DefaultValue: 0.,
			},
			output: 42.,
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": []string{}},
				Key:          "value",
				DefaultValue: 0.,
			},
			output: 0.,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Float64(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if got != tt.output {
				t.Errorf(
					"expected Float64(%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output,
					got,
				)
			}
		})
	}
}
