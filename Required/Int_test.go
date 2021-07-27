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
		Value         int
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
				Origin:       map[string]interface{}{"value": 1},
				Key:          "NonExisting",
				DefaultValue: 0,
			},
			output: outputStruct{
				Value:         0,
				IsValid:       false,
				MissingFields: []string{"NonExisting"},
			},
		},
		{
			name: "CaseString",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "1"},
				Key:          "value",
				DefaultValue: 0,
			},
			output: outputStruct{
				Value:         1,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseStringErr",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "one"},
				Key:          "value",
				DefaultValue: 0,
			},
			output: outputStruct{
				Value:         0,
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
		{
			name: "CaseInt",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 1},
				Key:          "value",
				DefaultValue: 0,
			},
			output: outputStruct{
				Value:         1,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseInt32",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": int32(1)},
				Key:          "value",
				DefaultValue: 0,
			},
			output: outputStruct{
				Value:         1,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseInt64",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": int64(1)},
				Key:          "value",
				DefaultValue: 0,
			},
			output: outputStruct{
				Value:         1,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseFloat64",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": float64(1)},
				Key:          "value",
				DefaultValue: 0,
			},
			output: outputStruct{
				Value:         1,
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": []string{"1"}},
				Key:          "value",
				DefaultValue: 0,
			},
			output: outputStruct{
				Value:         0,
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var missingFieldsGot []string
			gotValue, gotValid := Int(tt.input.Origin, tt.input.Key, &missingFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(gotValue) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on Int(%v,%v,%v,%v) = %v; got %v",
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
					"expected isValid on Int(%v,%v,%v,%v) = %v; got %v",
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
					"expected missingFields on Int(%v,%v,%v,%v) = %v; got %v",
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
