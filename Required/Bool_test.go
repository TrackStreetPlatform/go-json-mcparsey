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
		Value         bool
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
				Origin:       map[string]interface{}{"value": true},
				Key:          "NonExisting",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:         false,
				IsValid:       false,
				MissingFields: []string{"NonExisting"},
			},
		},
		{
			name: "CaseString",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "true"},
				Key:          "value",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:         true,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseStringErr",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "one"},
				Key:          "value",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:         false,
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
		{
			name: "CaseInt",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 1},
				Key:          "value",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:         true,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseFloat64",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": float64(1)},
				Key:          "value",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:         true,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseBool",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": true},
				Key:          "value",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:         true,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": []string{"1"}},
				Key:          "value",
				DefaultValue: false,
			},
			output: outputStruct{
				Value:         false,
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var missingFieldsGot []string
			gotValue, gotValid := Bool(tt.input.Origin, tt.input.Key, &missingFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(gotValue) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on Bool(%v,%v,%v,%v) = %v; got %v",
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
					"expected isValid on Bool(%v,%v,%v,%v) = %v; got %v",
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
					"expected missingFields on Bool(%v,%v,%v,%v) = %v; got %v",
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
