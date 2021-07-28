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
		Value         []string
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
				Origin:       map[string]interface{}{"value": "1"},
				Key:          "NonExisting",
				DefaultValue: []string{},
			},
			output: outputStruct{
				Value:         []string{},
				IsValid:       false,
				MissingFields: []string{"NonExisting"},
			},
		},
		{
			name: "CaseString",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "1,2"},
				Key:          "value",
				DefaultValue: []string{},
			},
			output: outputStruct{
				Value:         []string{"1", "2"},
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseInterfaceList",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": []interface{}{"1", "2"}},
				Key:          "value",
				DefaultValue: []string{}},
			output: outputStruct{
				Value:         []string{"1", "2"},
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "CaseInterfaceListWrongType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": []interface{}{"1", 42}},
				Key:          "value",
				DefaultValue: []string{},
			},
			output: outputStruct{
				Value:         []string{},
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
		{
			name: "CaseStringList",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": []string{"1", "2"}},
				Key:          "value",
				DefaultValue: []string{},
			},
			output: outputStruct{
				Value:         []string{"1", "2"},
				IsValid:       true,
				MissingFields: []string{""},
			},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": true},
				Key:          "value",
				DefaultValue: []string{},
			},
			output: outputStruct{
				Value:         []string{},
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var missingFieldsGot []string
			gotValue, gotValid := ArrayString(tt.input.Origin, tt.input.Key, &missingFieldsGot, tt.input.DefaultValue)
			if fmt.Sprint(gotValue) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on ArrayString(%v,%v,%v,%v) = %v; got %v",
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
					"expected isValid on ArrayString(%v,%v,%v,%v) = %v; got %v",
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
					"expected missingFields on ArrayString(%v,%v,%v,%v) = %v; got %v",
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
