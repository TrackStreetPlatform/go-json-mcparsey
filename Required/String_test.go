package Required

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue string
	}

	type outputStruct struct {
		Value          string
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
				DefaultValue: "default",
			},
			output: outputStruct{
				Value:          "default",
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "ValidValue",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "value",
				},
				Key:          "key",
				DefaultValue: "default",
			},
			output: outputStruct{
				Value:          "value",
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": false,
				},
				Key:          "key",
				DefaultValue: "default",
			},
			output: outputStruct{
				Value:          "default",
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requiredFieldsGot []string
			valueGot, isValidGot := String(tt.input.Origin, tt.input.Key, &requiredFieldsGot, tt.input.DefaultValue)
			if valueGot != tt.output.Value {
				t.Errorf(
					"expected value on String(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if isValidGot != tt.output.IsValid {
				t.Errorf(
					"expected isValid on String(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.IsValid,
					isValidGot,
				)
			}
			if fmt.Sprint(requiredFieldsGot) != fmt.Sprint(tt.output.RequiredFields) {
				t.Errorf(
					"expected requiredFields on String(%v,%v,requiredFields,%v) = %v; got %v",
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
