package Required

import (
	"fmt"
	"testing"
)

func TestFloat64(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue float64
	}

	type outputStruct struct {
		Value          float64
		IsValid        bool
		RequiredFields []string
	}

	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "NoKey",
			input: inputStruct{
				Origin:       map[string]interface{}{},
				Key:          "key",
				DefaultValue: -1.0,
			},
			output: outputStruct{
				Value:   -1.0,
				IsValid: false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "Int",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": 17,
				},
				Key:          "key",
				DefaultValue: -1.0,
			},
			output: outputStruct{
				Value:          17.0,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "Float",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": 8.3,
				},
				Key:          "key",
				DefaultValue: -1.0,
			},
			output: outputStruct{
				Value:          8.3,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "String",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "-163.12",
				},
				Key:          "key",
				DefaultValue: -1.0,
			},
			output: outputStruct{
				Value:          -163.12,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "invalid string",
				},
				Key:          "key",
				DefaultValue: -1.0,
			},
			output: outputStruct{
				Value:          -1.0,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "InvalidType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": []int{1, 2, 3},
				},
				Key:          "key",
				DefaultValue: -1.0,
			},
			output: outputStruct{
				Value:          -1.0,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requiredFieldsGot []string
			valueGot, isValidGot := Float64(tt.input.Origin, tt.input.Key, &requiredFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(valueGot) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on Float64(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if isValidGot != tt.output.IsValid {
				t.Errorf(
					"expected isValid on Float64(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.IsValid,
					isValidGot,
				)
			}

			if fmt.Sprint(requiredFieldsGot) != fmt.Sprint(tt.output.RequiredFields) {
				t.Errorf(
					"expected requiredFields on Float64(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.RequiredFields,
					requiredFieldsGot,
				)
			}

		})
	}
}
