package Required

import (
	"fmt"
	"testing"
)

func TestInterfaceArray(t *testing.T) {
	type inputStruct struct {
		Origin map[string]interface{}
		Key    string
	}

	type outputStruct struct {
		Value          []map[string]interface{}
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
				Origin: map[string]interface{}{},
				Key:    "key",
			},
			output: outputStruct{
				Value:          []map[string]interface{}{},
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "EmptyInput",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": make([]map[string]interface{}, 0),
				},
				Key: "key",
			},
			output: outputStruct{
				Value:          []map[string]interface{}{},
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "ValidInput",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": []interface{}{
						map[string]interface{}{"one": 1, "two": 2.0},
					},
				},
				Key: "key",
			},
			output: outputStruct{
				Value:          []map[string]interface{}{{"one": 1, "two": 2.0}},
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "ExtraArgs",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": []interface{}{
						false,
						map[string]interface{}{"one": 1, "two": 2.0},
						1,
					},
				},
				Key: "key",
			},
			output: outputStruct{
				Value:          []map[string]interface{}{{"one": 1, "two": 2.0}},
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "InvalidType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": true,
				},
				Key: "key",
			},
			output: outputStruct{
				Value:          []map[string]interface{}{},
				IsValid:        false,
				RequiredFields: []string{"key"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requiredFieldsGot []string
			valueGot, isValidGot := InterfaceArray(tt.input.Origin, tt.input.Key, &requiredFieldsGot)
			if fmt.Sprint(valueGot) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on InterfaceArray(%v,%v,requiredFields) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.output.Value,
					valueGot,
				)
			}
			if isValidGot != tt.output.IsValid {
				t.Errorf(
					"expected isValid on InterfaceArray(%v,%v,requiredFields) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.output.IsValid,
					isValidGot,
				)
			}
			if fmt.Sprint(requiredFieldsGot) != fmt.Sprint(tt.output.RequiredFields) {
				t.Errorf(
					"expected requiredFields on InterfaceArray(%v,%v,requiredFields) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.output.RequiredFields,
					requiredFieldsGot,
				)
			}
		})
	}
}
