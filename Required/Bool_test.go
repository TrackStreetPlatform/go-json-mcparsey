package Required

import (
	"fmt"
	"testing"
)

func TestBool(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue bool
	}

	type outputStruct struct {
		Value          bool
		IsValid        bool
		RequiredFields []string
	}

	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "NoKeyDefaultTrue",
			input: inputStruct{
				Origin:       map[string]interface{}{},
				Key:          "key",
				DefaultValue: true,
			},
			output: outputStruct{
				Value:   true,
				IsValid: false,
				RequiredFields: []string{
					"key",
				},
			},
		}, {
			name: "NoKeyDefaultFalse",
			input: inputStruct{
				Origin:       map[string]interface{}{},
				Key:          "key",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:   false,
				IsValid: false,
				RequiredFields: []string{
					"key",
				},
			},
		},
		{
			name: "TrueString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "TRUE",
				},
				Key:          "key",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:          true,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "FalseString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "0",
				},
				Key:          "key",
				DefaultValue: true,
			},
			output: outputStruct{
				Value:          false,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "not a valid string",
				},
				Key:          "key",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:   false,
				IsValid: false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "TrueInt",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": 1,
				},
				Key:          "key",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:          true,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "FalseInt",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": 0,
				},
				Key:          "key",
				DefaultValue: true,
			},
			output: outputStruct{
				Value:          false,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "TrueFloat",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": 1.0,
				},
				Key:          "key",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:          true,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "FalseFloat",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": 0.0,
				},
				Key:          "key",
				DefaultValue: true,
			},
			output: outputStruct{
				Value:          false,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "Bool",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": true,
				},
				Key:          "key",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:          true,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": []int{1, 2, 3},
				},
				Key:          "key",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:   false,
				IsValid: false,
				RequiredFields: []string{"key"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requiredFieldsGot []string
			valueGot, isValidGot := Bool(tt.input.Origin, tt.input.Key, &requiredFieldsGot, tt.input.DefaultValue)
			if valueGot != tt.output.Value {
				t.Errorf(
					"expected value on Bool(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if isValidGot != tt.output.IsValid {
				t.Errorf(
					"expected isValid on Bool(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.IsValid,
					isValidGot,
				)
			}
			if fmt.Sprint(requiredFieldsGot) != fmt.Sprint(tt.output.RequiredFields) {
				t.Errorf(
					"expected requiredFields on Bool(%v,%v,requiredFields,%v) = %v; got %v",
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
