package Optional

import (
	"reflect"
	"testing"
)

func TestMapStringInterface(t *testing.T) {
	type inputStruct struct {
		Origin       map[string]interface{}
		Key          string
		DefaultValue map[string]interface{}
	}
	tests := []struct {
		name   string
		input  inputStruct
		output map[string]interface{}
	}{
		{
			name: "NonExistingKey",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": map[string]interface{}{"attribute": 42},
				}, Key: "NonExisting", DefaultValue: map[string]interface{}{}},
			output: map[string]interface{}{},
		},
		{
			name: "CaseMapStringInterface",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": map[string]interface{}{"attribute": 42},
				}, Key: "value", DefaultValue: map[string]interface{}{}},
			output: map[string]interface{}{"attribute": 42},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": 42,
				}, Key: "value", DefaultValue: map[string]interface{}{}},
			output: map[string]interface{}{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := MapStringInterface(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf(
					"expected MapStringInterface(%v,%v,%v) = %v; got %v",
					tt.input.Origin,
					tt.input.Key,
					tt.input.DefaultValue,
					tt.output,
					got,
				)
			}
		})
	}
}
