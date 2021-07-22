package Required

import (
	"fmt"
	"testing"
)

func TestArrayString(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue []string
	}

	type outputStruct struct {
		Value          []string
		IsValid        bool
		RequiredFields []string
	}

	defaultValue := []string{"default", "value"}
	oneTwoThree := []string{"one", "two", "three"}

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
				Value:   defaultValue,
				IsValid: false,
				RequiredFields: []string{"key"},
			},
		},
		{
			name: "CSVString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": "one,two,three",
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          oneTwoThree,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "StringArray",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": []string{"one", "two", "three"},
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          oneTwoThree,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "StringInterfaceArray",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": []interface{}{"one", "two", "three"},
				},
				Key:          "key",
				DefaultValue: defaultValue,
			},
			output: outputStruct{
				Value:          oneTwoThree,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
		{
			name: "IntArray",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": []interface{}{"one", 2, 3},
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
		{
			name: "InvalidType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"key": 3,
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
			valueGot, isValidGot := ArrayString(tt.input.Origin, tt.input.Key, &requiredFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(valueGot) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on ArrayString(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if isValidGot != tt.output.IsValid {
				t.Errorf(
					"expected isValid on ArrayString(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.IsValid,
					isValidGot,
				)
			}
			if fmt.Sprint(requiredFieldsGot) != fmt.Sprint(tt.output.RequiredFields) {
				t.Errorf(
					"expected requiredFields on ArrayString(%v,%v,requiredFields,%v) = %v; got %v",
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
