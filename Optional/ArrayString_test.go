package Optional

import (
	"fmt"
	"testing"
)

type inputStruct struct {
	Origin       map[string]interface{}
	Key          string
	DefaultValue []string
}

type testArrayString struct {
	name   string
	input  inputStruct
	output []string
}

func TestArrayString(t *testing.T) {
	tests := []testArrayString{
		{
			name: "NonExistentKeys",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "1"},
				Key:          "NonExisting",
				DefaultValue: []string{},
			},
			output: []string{},
		},
		{
			name: "CaseString",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": "1,2"},
				Key:          "value",
				DefaultValue: []string{}},
			output: []string{"1", "2"},
		},
		{
			name: "UnsupportedType",
			input: inputStruct{
				Origin:       map[string]interface{}{"value": 42},
				Key:          "value",
				DefaultValue: []string{}},
			output: []string{},
		},
		{
			name: "InterfaceArrayProper",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []interface{}{"disused", "sdfsd"},
				},
				Key:          "value",
				DefaultValue: []string{},
			},
			output: []string{"disused", "sdfsd"},
		},
		{
			name: "InterfaceArrayProperOneItemWrongType",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []interface{}{"disused", "sdfsd", 42},
				},
				Key:          "value",
				DefaultValue: []string{},
			},
			output: []string{},
		},
		{
			name: "StringArrayProper",
			input: inputStruct{
				Origin: map[string]interface{}{
					"value": []string{"disused", "sdfsd"},
				},
				Key:          "value",
				DefaultValue: []string{},
			},
			output: []string{"disused", "sdfsd"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrayString(tt.input.Origin, tt.input.Key, tt.input.DefaultValue)
			if fmt.Sprint(got) != fmt.Sprint(tt.output) {
				t.Errorf(
					"expected ArrayString(%v,%v,%v) = %v; got %v",
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
