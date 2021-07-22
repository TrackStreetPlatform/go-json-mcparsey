package Required

import (
	"fmt"
	"testing"
)

func TestInt(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue int
	}

	type outputStruct struct {
		Value          int
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
				DefaultValue: -1,
			},
			output: outputStruct{
				Value:   -1,
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
				DefaultValue: -1,
			},
			output: outputStruct{
				Value:          17,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "Int32",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": int32(2147483647),
				},
				Key:          "key",
				DefaultValue: -1,
			},
			output: outputStruct{
				Value:          2147483647,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "Int64",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": int64(-9223372036854775808),
				},
				Key:          "key",
				DefaultValue: -1,
			},
			output: outputStruct{
				Value:          -9223372036854775808,
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
				DefaultValue: -1,
			},
			output: outputStruct{
				Value:          8,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "String",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "-163",
				},
				Key:          "key",
				DefaultValue: -1,
			},
			output: outputStruct{
				Value:          -163,
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
				DefaultValue: -1,
			},
			output: outputStruct{
				Value:          -1,
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
				DefaultValue: -1,
			},
			output: outputStruct{
				Value:          -1,
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requiredFieldsGot []string
			valueGot, isValidGot := Int(tt.input.Origin, tt.input.Key, &requiredFieldsGot, tt.input.DefaultValue)
			if valueGot != tt.output.Value {
				t.Errorf(
					"expected value on Int(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if isValidGot != tt.output.IsValid {
				t.Errorf(
					"expected isValid on Int(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.IsValid,
					isValidGot,
				)
			}
			if fmt.Sprint(requiredFieldsGot) != fmt.Sprint(tt.output.RequiredFields) {
				t.Errorf(
					"expected requiredFields on Int(%v,%v,requiredFields,%v) = %v; got %v",
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
