package Required

import (
	"fmt"
	"testing"
)

func TestMapStringInterface(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue map[string]interface{}
	}

	type outputStruct struct {
		Value          map[string]interface{}
		IsValid        bool
		RequiredFields []string
	}

	defaultValue := map[string]interface{}{"default": "value"}
	correctValue := map[string]interface{}{"one": 1, "two": 2.0}

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
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          defaultValue,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "ValidValue",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": correctValue,
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          correctValue,
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
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          defaultValue,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requiredFieldsGot []string
			valueGot, isValidGot := MapStringInterface(tt.input.Origin, tt.input.Key, &requiredFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(valueGot) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on MapStringInterface(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if fmt.Sprint(isValidGot) != fmt.Sprint(tt.output.IsValid) {
				t.Errorf(
					"expected isValid on MapStringInterface(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.IsValid,
					isValidGot,
				)
			}

			if fmt.Sprint(requiredFieldsGot) != fmt.Sprint(tt.output.RequiredFields) {
				t.Errorf(
					"expected requiredFields on MapStringInterface(%v,%v,requiredFields,%v) = %v; got %v",
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
