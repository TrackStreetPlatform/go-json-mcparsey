package Required

import (
	"fmt"
	"testing"
)

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

func TestBool(t *testing.T) {
	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "NoKeyDefaultTrue",
			input: inputStruct{
				Origin:       map[string]interface{}{},
				Key:          "use_tor",
				DefaultValue: true,
			},
			output: outputStruct{
				Value:   true,
				IsValid: false,
				RequiredFields: []string{
					"use_tor",
				}},
		}, {
			name: "NoKeyDefaultFalse",
			input: inputStruct{
				Origin:       map[string]interface{}{},
				Key:          "use_tor",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:   false,
				IsValid: false,
				RequiredFields: []string{
					"use_tor",
				}},
		},
		{
			name: "TrueString",
			input: inputStruct{
				Origin: map[string]interface{}{
					"use_tor": "TRUE",
				},
				Key:          "use_tor",
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
					"use_tor": "0",
				},
				Key:          "use_tor",
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
					"use_tor": "not a valid string",
				},
				Key:          "use_tor",
				DefaultValue: true,
			},
			output: outputStruct{
				Value:   false,
				IsValid: true,
				RequiredFields: []string{
					"use_tor",
				},
			},
		},
		{
			name: "TrueInt",
			input: inputStruct{
				Origin: map[string]interface{}{
					"use_tor": 1,
				},
				Key:          "use_tor",
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
					"use_tor": 0,
				},
				Key:          "use_tor",
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
					"use_tor": 1.0,
				},
				Key:          "use_tor",
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
					"use_tor": 0.0,
				},
				Key:          "use_tor",
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
					"use_tor": true,
				},
				Key:          "use_tor",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:          true,
				IsValid:        true,
				RequiredFields: []string{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var requiredFieldsGot []string
			valueGot, isValidGot := Bool(tt.input.Origin, tt.input.Key, &requiredFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(valueGot) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on Bool(%v,%v,requiredFields,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output.Value,
					valueGot,
				)
			}
			if fmt.Sprint(isValidGot) != fmt.Sprint(tt.output.IsValid) {
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
