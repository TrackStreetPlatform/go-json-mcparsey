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
		Value         []map[string]interface{}
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
				Origin: map[string]interface{}{"value": 1.},
				Key:    "NonExisting",
			},
			output: outputStruct{
				Value:         []map[string]interface{}{},
				IsValid:       false,
				MissingFields: []string{"NonExisting"},
			},
		},
		{
			name: "CommonCase",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []interface{}{map[string]interface{}{"test": 42}},
				},
				Key: "value",
			},
			output: outputStruct{
				Value:         []map[string]interface{}{{"test": 42}},
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "ListInterfaceMultipleEntries",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []interface{}{
						map[string]interface{}{"test1": 42},
						map[string]interface{}{"test2": "42"}},
				},
				Key: "value",
			},
			output: outputStruct{
				Value:         []map[string]interface{}{{"test1": 42}, {"test2": "42"}},
				IsValid:       true,
				MissingFields: []string{},
			},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin: map[string]interface{}{"value": []string{"1"}},
				Key:    "value",
			},
			output: outputStruct{
				Value:         []map[string]interface{}{},
				IsValid:       false,
				MissingFields: []string{"value"},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var missingFieldsGot []string
			gotValue, gotValid := InterfaceArray(tt.input.Origin, tt.input.Key, &missingFieldsGot)
			if fmt.Sprint(gotValue) != fmt.Sprint(tt.output.Value) {
				t.Errorf(
					"expected value on InterfaceArray(%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.output.Value,
					gotValue,
				)
			}
			if fmt.Sprint(gotValid) != fmt.Sprint(tt.output.IsValid) {
				t.Errorf(
					"expected isValid on InterfaceArray(%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.output.IsValid,
					gotValid,
				)
			}
			if fmt.Sprint(missingFieldsGot) != fmt.Sprint(tt.output.MissingFields) {
				t.Errorf(
					"expected missingFields on InterfaceArray(%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					missingFieldsGot,
					tt.output.MissingFields,
					missingFieldsGot,
				)
			}
		})
	}
}
