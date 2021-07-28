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
		Value         map[string]interface{}
		IsValid       bool
		MissingFields []string
	}
	tests := []struct {
		name   string
		input  inputStruct
		output outputStruct
	}{
		{
			name: "NonExistentKeys",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 42},
				Key:          "NonExisting",
				DefaultValue: map[string]interface{}{},
			},
			output: outputStruct{
				Value:         map[string]interface{}{},
				IsValid:       false,
				MissingFields: []string{"NonExisting"},
			},
		},
		{
			name: "CaseMapStringInterface",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": map[string]interface{}{"test": 42},
				},
				Key:          "value",
				DefaultValue: map[string]interface{}{},
			},
			output: outputStruct{
				Value:         map[string]interface{}{"test": 42},
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": []string{"1"}},
				Key:          "value",
				DefaultValue: map[string]interface{}{},
			},
			output: outputStruct{
				Value:         map[string]interface{}{},
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var missingFieldsGot []string
			gotValue, gotValid := MapStringInterface(tt.input.Origin, tt.input.Key, &missingFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(gotValue) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on MapStringInterface(%v,%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.input.DefaultValue,
					tt.output.Value,
					gotValue,
				)
			}
			if fmt.Sprint(gotValid) != fmt.Sprint(tt.output.IsValid) {
				t.Errorf(
					"expected isValid on MapStringInterface(%v,%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.input.DefaultValue,
					tt.output.IsValid,
					gotValid,
				)
			}
			if fmt.Sprint(missingFieldsGot) != fmt.Sprint(tt.output.MissingFields) {
				t.Errorf(
					"expected missingFields on MapStringInterface(%v,%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.input.DefaultValue,
					tt.output.MissingFields,
					missingFieldsGot,
				)
			}
		})
	}
}
